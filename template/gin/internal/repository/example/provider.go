package example

import (
	"boilerplate/config"
	"boilerplate/config/infra"
	"boilerplate/internal/model"
	"context"
	"github.com/saucon/sauron/v2/pkg/log"
)

type IBankRepo interface {
	ResolveByFilter(ctx context.Context, filter model.Filter) (model.BankList, error)
}

type BankRepo struct {
	log   *log.LogCustom
	cfg   *config.Config
	Infra *infra.Infra
}

func ProvideBankRepo(log *log.LogCustom, cfg *config.Config, infra *infra.Infra) IBankRepo {
	return &BankRepo{
		log:   log,
		cfg:   cfg,
		Infra: infra,
	}
}
