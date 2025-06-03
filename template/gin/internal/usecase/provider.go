package usecase

import (
	"boilerplate/config"
	"boilerplate/config/infra"
	"boilerplate/external"
	"boilerplate/internal/repository/example"
	"boilerplate/internal/repository/tx"
	"github.com/saucon/sauron/v2/pkg/log"
)

type IUsecase interface {
}

type usecase struct {
	log         *log.LogCustom
	cfg         *config.Config
	redisClient *infra.Redis
	ext         *external.External
	txManager   tx.TxManager
	exampleRepo example.IBankRepo
}

func ProvideUsc(log *log.LogCustom, cfg *config.Config,
	redisClient *infra.Redis,
	ext *external.External,
	txManager tx.TxManager,
	exampleRepo example.IBankRepo,
) IUsecase {
	return &usecase{
		log:         log,
		cfg:         cfg,
		redisClient: redisClient,
		ext:         ext,
		txManager:   txManager,
		exampleRepo: exampleRepo,
	}
}
