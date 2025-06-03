package infra

import (
	"boilerplate/config"
	"boilerplate/utils"
	"fmt"
	saurondb "github.com/saucon/sauron/v2/pkg/db"
	"github.com/saucon/sauron/v2/pkg/log"
	"time"
)

func ProvideDbPsql(cfg *config.Config, logger *log.LogCustom) (dbRes *saurondb.Database) {
	defer func() {
		logger.Info(log.LogData{
			Err:         dbRes.DB.Error,
			Description: fmt.Sprintf("connect db with %v using %v", cfg.AppEnvMode.Mode, cfg.DBConfig.DBPostgresConfig["postgres"]),
			StartTime:   time.Now(),
			Response:    nil,
		})
	}()
	switch cfg.AppEnvMode.Mode {
	case utils.DEV:
		dbRes = saurondb.NewDB(&cfg.DBConfig, logger, "main_db", "", false, "postgres")
		return
	case utils.DEV_TEST:
		dbRes = saurondb.NewDB(&cfg.DBConfig, logger, "test_db", "", false, "postgres")
		return
	case utils.PROD:
		dbRes = saurondb.NewDB(&cfg.DBConfig, logger, "main_db", "", false, "postgres")
		return
	default:
		dbRes = saurondb.NewDB(&cfg.DBConfig, logger, "main_db", "", false, "postgres")
		return
	}
}
