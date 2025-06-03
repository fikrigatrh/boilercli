package traceheaderid

import (
	"boilerplate/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TraceIDHandler struct {
}

func New(r *gin.RouterGroup) {
	handler := &TraceIDHandler{}

	r.Use(handler.traceIDHandler)
}

func (eh *TraceIDHandler) traceIDHandler(c *gin.Context) {
	id := uuid.New().String()
	c.Request.Header.Set("X-Service-Req-ID", id)

	utils.TraceHeaderWrapGinContext(c).Next()
}
