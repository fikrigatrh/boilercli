package middleware

import (
	"boilerplate/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (p *PaymentMiddleware) Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(utils.HeaderRequestID)
		timestamp := c.GetHeader(utils.TimestampHeader)
		signature := c.GetHeader(utils.SignatureHeader)

		// Optional: Validate required headers
		if requestID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing Request-ID header"})
			return
		}

		if timestamp == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing Timestamp header"})
			return
		}

		if signature == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing Signature header"})
			return
		}

		// You can set in context so it's accessible in handler
		c.Set(utils.ContextRequestIDKey, requestID)

		c.Next()
	}
}

func (p *PaymentMiddleware) HeaderUseFiber() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(utils.HeaderRequestID)
		timestamp := c.Get(utils.TimestampHeader)
		signature := c.Get(utils.SignatureHeader)
		clientId := c.Get(utils.ClientKeySignature)

		// Validate headers
		if requestID == "" || len(requestID) > 36 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"responseCode":    "4001101",
				"responseMessage": "Missing or Invalid Format Request-ID header",
			})
		}

		if timestamp == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"responseCode":    "4001102",
				"responseMessage": "Missing or Invalid Format Timestamp header",
			})
		}

		if signature == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"responseCode":    "4001103",
				"responseMessage": "Missing or Invalid Format Signature header",
			})
		}

		if clientId == "" || len(clientId) > 36 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"responseCode":    "4001104",
				"responseMessage": "Missing or Invalid Format Client-Key header",
			})
		}

		// Save to locals (similar to Gin's c.Set)
		c.Locals(utils.ContextRequestIDKey, requestID)
		c.Locals(utils.ClientKeySignature, clientId)

		traceHeader := map[string]string{
			"trace_srvc_id": requestID,
			"signature":     signature,
			"client_id":     clientId,
		}

		// Store it in the context's locals
		c.Locals(utils.TraceHeaderKey, traceHeader)

		// Continue to next handler
		return c.Next()
	}
}
