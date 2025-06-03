package router

import (
	"boilerplate/internal/middleware"
	"boilerplate/utils"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

func (h *Route) SetupRoute(app *fiber.App) *fiber.App {
	app.Use(fiberLog.New(fiberLog.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Define the health check endpoint
	// Register the health check endpoint
	app.Get("/health-check", healthCheckUseFiber)
	app.Post("/create-signature", createComponentSignature)

	app.Use(fiberCors.New(fiberCors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTION, PATCH",
	}))

	api := app.Group("/v1/example")

	secureApi := api.Group("/")
	middleware.ProviderMiddleware(secureApi, h.Cfg, h.Log, h.usecase)

	return app
}

// @title Health Check Service API
// @description API for checking the health status of the service.

// @Summary Health Check
// @Description Returns the health status of the service
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func healthCheckUseFiber(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "UP",
		"message": "Service is running",
	})
}

func processTimeout(handler fiber.Handler, timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Set the context to the Fiber context
		c.SetUserContext(ctx)

		// Create a channel to signal handler completion
		processDone := make(chan struct{}, 1)
		var handlerErr error

		// Run the handler in a goroutine
		go func() {
			defer func() {
				if r := recover(); r != nil {
					// Handle any panic in the handler to prevent crashes
					handlerErr = fiber.NewError(fiber.StatusInternalServerError, "Handler panicked")
				}
				close(processDone)
			}()
			// Execute the handler and capture any error
			handlerErr = handler(c)
		}()

		// Wait for either the handler to complete or the context to timeout
		select {
		case <-ctx.Done():
			// Context timed out or was canceled
			return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{
				"responseCode":    "4080100",
				"responseMessage": "Request Process Timeout",
			})
		case <-processDone:
			// Handler completed, return its error (if any)
			return handlerErr
		}
	}
}

// @Summary Create Signature
// @Description Returns the signature
// @Tags Signature
// @Param request body signatureRequest true "request for create signature"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /create-signature [POST]
func createComponentSignature(c *fiber.Ctx) error {
	var req signatureRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	xTimestamp := time.Now().Format(utils.XTimestampLayoutFormat)
	stringToSign := req.Method + ":" + req.Endpoint + ":" + xTimestamp

	fmt.Println(stringToSign)
	hashed := utils.Hash([]byte(stringToSign))

	// Sign the hash using the client secret.
	signature, err := utils.SignWithECC(req.PrivateKey, hashed)
	if err != nil {
		panic(err)
	}

	// Encode the signature in base64 for transmission or storage.
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"timestamp":       xTimestamp,
		"signatureBase64": signatureBase64,
	})
}

type signatureRequest struct {
	PrivateKey string `json:"privateKey"`
	Method     string `json:"method"`
	Endpoint   string `json:"endpoint"`
}
