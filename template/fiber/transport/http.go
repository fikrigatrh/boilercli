package transport

import (
	"boilerplate/config"
	"boilerplate/config/router"
	"boilerplate/docs"
	_ "boilerplate/docs"
	"boilerplate/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gookit/color"
	"github.com/rs/zerolog/log"
	sauconLog "github.com/saucon/sauron/v2/pkg/log"
	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
	defaultLog "log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// ServerState is an indicator if this server's state.
type ServerState int

const (
	// ServerStateReady indicates that the server is ready to serve.
	ServerStateReady ServerState = iota + 1
	// ServerStateInGracePeriod indicates that the server is in its grace
	// period and will shut down after it is done cleaning up.
	ServerStateInGracePeriod
	// ServerStateInCleanupPeriod indicates that the server no longer
	// responds to any requests, is cleaning up its internal state, and
	// will shut down shortly.
	ServerStateInCleanupPeriod
)

// HTTP is the HTTP server.
type HTTP struct {
	Config      *config.Config
	Route       router.Route
	State       ServerState
	FiberServer *fiber.App
	Log         *sauconLog.LogCustom
}

func ProvideHttp(Config *config.Config,
	route router.Route,
	log *sauconLog.LogCustom,
) *HTTP {
	srv := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Log the panic error and stack trace
			logError(err)
			// Return structured error response
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"responseCode":    "00000",
				"responseMessage": err.Error(),
			})
		},
	})

	srv.Use(fiberLogger.New())
	srv.Use(fiberRecover.New())

	return &HTTP{
		Config:      Config,
		Route:       route,
		FiberServer: srv,
		Log:         log,
	}
}

func (h *HTTP) Serve() {
	app := h.Route.SetupRoute(h.FiberServer)

	// Log all registered routes
	for _, routes := range app.Stack() {
		for _, route := range routes {
			if strings.Contains(route.Method, "HEAD") {
				continue
			}
			log.Printf("[FIBER-debug] %s\t%-20s \n", route.Method, route.Path)
		}
	}

	h.setupSwaggerDocs(app)

	// Catch-all for undefined routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"responseCode":    "4040000",
			"responseMessage": "Route Not Found",
		})
	})

	h.shutdownApplication(app)

	h.startApplication(app, h.Config.EnvConfig.AppConfig.Port)

	h.cleanUpApplication()
}

func (h *HTTP) setupGracefulShutdown() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go h.respondToSigterm(done)
}

func (h *HTTP) respondToSigterm(done chan os.Signal) {
	<-done
	defer os.Exit(0)

	shutdownConfig := h.Config.Server.Shutdown

	log.Info().Msg("Received SIGTERM.")
	log.Info().Int64("seconds", shutdownConfig.GracePeriodSeconds).Msg("Entering grace period.")
	h.State = ServerStateInGracePeriod
	time.Sleep(time.Duration(shutdownConfig.GracePeriodSeconds) * time.Second)

	log.Info().Int64("seconds", shutdownConfig.CleanupPeriodSeconds).Msg("Entering cleanup period.")
	h.State = ServerStateInCleanupPeriod
	time.Sleep(time.Duration(shutdownConfig.CleanupPeriodSeconds) * time.Second)

	log.Info().Msg("Cleaning up completed. Shutting down now.")
}

func (h *HTTP) startApplication(r *fiber.App, port string) {

	// Start the server
	err := r.Listen(":" + port)
	if err != nil {
		panic(err)
	}

}

func (h *HTTP) shutdownApplication(r *fiber.App) {

	// Implement graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-ctx.Done()
		defaultLog.Println("🛑 Shutting down...")
		_ = r.Shutdown()
		stop()
	}()
}

func (h *HTTP) cleanUpApplication() {
	shutdownConfig := h.Config.Server.Shutdown
	defaultLog.Println("Running cleanup tasks...")
	// wait n seconds for the server to shutdown
	time.Sleep(time.Duration(shutdownConfig.CleanupPeriodSeconds))
	context.WithTimeout(context.Background(), time.Duration(shutdownConfig.GracePeriodSeconds))
	//DBCloseConnection()
	defaultLog.Println("✅ Graceful shutdown complete")
}

func (h *HTTP) setupSwaggerDocs(app *fiber.App) {
	if h.Config.AppEnvMode.Mode == utils.DEV || h.Config.AppEnvMode.Mode == utils.DEV_TEST {
		// Set Swagger Info
		docs.SwaggerInfo.Title = h.Config.EnvConfig.AppConfig.Name
		docs.SwaggerInfo.Version = h.Config.EnvConfig.AppConfig.Version

		// Setup Swagger
		app.Get("/swagger/*", fiberSwagger.WrapHandler)

		log.Info().Fields(map[string]interface{}{
			"url": h.Config.EnvConfig.AppConfig.Host + ":" + h.Config.EnvConfig.AppConfig.Port,
		}).Msg("Swagger documentation enabled.")
	}
}

func logError(err error) {
	color.Red.Printf("[PANIC] %v\n", err)
}
