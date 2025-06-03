package infra

import (
	"boilerplate/config"
	saurondb "github.com/saucon/sauron/v2/pkg/db"
	"github.com/saucon/sauron/v2/pkg/log"
)

type Infra struct {
	DbPsql *saurondb.Database
}

func ProvideInfra(cfg *config.Config, logV2 *log.LogCustom) *Infra {
	infra := &Infra{
		DbPsql: ProvideDbPsql(cfg, logV2),
	}

	if cfg.DBConfig.EnableAutoMigration {
		// if you want to migrate table
	}
	return infra
}
