package utils

import (
	bytes2 "bytes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

//
//func Failed(c *gin.Context, applies ...func(b *Base)) {
//	baseResponse := Failure()
//	for _, apply := range applies {
//		apply(baseResponse)
//	}
//
//	respond(c, baseResponse)
//}

// ErrorMessage ...
//func respond(c *gin.Context, baseResponse *Base) {
//	c.JSON(baseResponse.StatusCode, baseResponse)
//}

func Failed(c *fiber.Ctx, applies ...func(b *Base)) error {
	baseResponse := Failure()
	for _, apply := range applies {
		apply(baseResponse)
	}

	return respond(c, baseResponse)
}

func respond(c *fiber.Ctx, baseResponse *Base) error {
	return c.Status(baseResponse.StatusCode).JSON(baseResponse)
}

type Response[T any] struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Data            T      `json:"data,omitempty"`
}

func (r *Response[T]) SetToSuccess() {
	r.ResponseCode = "2000000"
	r.ResponseMessage = "Success"
}

func (r *Response[T]) SetToSuccessCreated() {
	r.ResponseCode = "2010000"
	r.ResponseMessage = "Success"
}

func TraceHeaderWrapGinContext(c *gin.Context) *gin.Context {
	c.Set("trace_header",
		map[string]string{
			"trace_srvc_id": c.Request.Header.Get("X-BRC-Service-Req-ID"),
			"trace_gw_id":   c.Request.Header.Get("X-BRC-Gateway-Req-ID"),
			"trace_ui_id":   c.Request.Header.Get("X-BRC-UI-Req-ID"),
		})

	return c
}

func PadLeft(str string, pad string, length int) string {
	var buffer bytes2.Buffer
	for i := 0; i < (length - len(str)); i = i + len(pad) {
		buffer.WriteString(pad)
	}
	result := buffer.String() + str
	return result[0:length]
}
