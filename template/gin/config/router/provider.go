package router

import (
	"boilerplate/config"
	"boilerplate/config/infra"
	"boilerplate/internal/handler"
	"boilerplate/internal/usecase"
	"github.com/saucon/sauron/v2/pkg/log"
)

// Route : populate all domain handler
type Route struct {
	Cfg         *config.Config
	redisClient *infra.Redis
	Log         *log.LogCustom
	handler     handler.Handler
	usecase     usecase.IUsecase
}

func ProvideRoute(
	cfg *config.Config,
	redisClient *infra.Redis,
	log *log.LogCustom,
	handler handler.Handler,
	usecase usecase.IUsecase,
) Route {
	return Route{
		Cfg:         cfg,
		redisClient: redisClient,
		Log:         log,
		handler:     handler,
		usecase:     usecase,
	}
}
