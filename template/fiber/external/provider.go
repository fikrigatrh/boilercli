package external

import (
	"boilerplate/config"
	"boilerplate/config/logger"
	"github.com/saucon/sauron/v2/pkg/log"
)

type External struct {
}

func ProvideExternalSvc(config *config.Config, logger *log.LogCustom,
	logDb *logger.LoggerDb) *External {
	return &External{}
}
