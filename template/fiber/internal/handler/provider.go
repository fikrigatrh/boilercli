package handler

import (
	"boilerplate/config"
	"boilerplate/internal/usecase"
	"github.com/saucon/sauron/v2/pkg/log"
)

type Handler struct {
	cfg     *config.Config
	log     *log.LogCustom
	usecase usecase.IUsecase
}

func ProvideHandler(
	cfg *config.Config,
	l *log.LogCustom,
	usecase usecase.IUsecase,
) Handler {
	return Handler{
		cfg:     cfg,
		log:     l,
		usecase: usecase,
	}
}
