package middleware

import (
	"boilerplate/config"
	"boilerplate/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/saucon/sauron/v2/pkg/log"
)

type IPaymentMiddleware interface {
	HeaderUseFiber() fiber.Handler
}

type PaymentMiddleware struct {
	Log    *log.LogCustom
	Config *config.Config
	usc    usecase.IUsecase
}

func ProviderMiddleware(app fiber.Router, c *config.Config,
	l *log.LogCustom,
	usc usecase.IUsecase,
) IPaymentMiddleware {
	mid := PaymentMiddleware{
		Log:    l,
		Config: c,
		usc:    usc,
	}
	app.Use(mid.HeaderUseFiber())

	return &mid
}
