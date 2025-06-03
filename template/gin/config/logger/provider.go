package logger

import (
	"boilerplate/config"
	"boilerplate/utils"
	"github.com/saucon/sauron/v2/pkg/log"
)

func ProvideLogger(cfg *config.Config) *log.LogCustom {
	cfg.LogConfig.AppConfig.Host = cfg.EnvConfig.AppConfig.Host
	cfg.LogConfig.AppConfig.Name = cfg.EnvConfig.AppConfig.Name
	cfg.LogConfig.AppConfig.Version = cfg.EnvConfig.AppConfig.Version
	cfg.LogConfig.AppConfig.Port = cfg.EnvConfig.AppConfig.Port

	switch cfg.AppEnvMode.Mode {
	case utils.DEV, utils.DEV_TEST:
		logger := log.NewLogCustom(&cfg.LogConfig)
		logger.PrettyPrintJSON(cfg.AppEnvMode.IsPrettyLog)
		return logger
	case utils.PROD:
		logger := log.NewLogCustom(&cfg.LogConfig)
		logger.PrettyPrintJSON(cfg.AppEnvMode.IsPrettyLog)
		return logger
	default:
		logger := log.NewLogCustom(&cfg.LogConfig)
		logger.PrettyPrintJSON(cfg.AppEnvMode.IsPrettyLog)
		return logger
	}
}
