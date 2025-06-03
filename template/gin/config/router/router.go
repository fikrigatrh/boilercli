package router

import (
	"boilerplate/internal/middleware/authmidware"
	"boilerplate/internal/middleware/traceheaderid"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (r *Route) SetupRoute(router *gin.Engine) {
	// Define the health check endpoint
	// Register the health check endpoint
	router.GET("/health-check", healthCheck)

	v1 := router.Group("v1")
	sample := v1.Group("/sample")

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = r.Cfg.Cors.AllowOrigins
	configCors.AllowHeaders = r.Cfg.Cors.AllowHeaders
	configCors.AllowMethods = r.Cfg.Cors.AllowMethods
	configCors.AllowCredentials = r.Cfg.Cors.AllowCredentials

	authmidware.New(sample, *r.Cfg, r.Log)
	traceheaderid.New(sample)
	sample.Use(cors.New(configCors))
}

// healthCheck returns the health status of the service.
// @Summary Health Check
// @Description Returns the health status of the service
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"message": "Service is running",
	})
}

func (r *Route) processTimeout(gh gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var duration time.Duration
		duration = time.Duration(5)
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration*time.Second)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		processDone := make(chan bool)
		go func() {
			gh(c)
			processDone <- true
		}()

		select {
		case <-ctx.Done():
			c.JSON(http.StatusRequestTimeout, gin.H{
				"responseCode":    "4080100",
				"responseMessage": "Request Process Timeout",
			})
		case <-processDone:
		}
	}
}
