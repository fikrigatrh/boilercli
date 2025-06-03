package config

import (
	"github.com/saucon/sauron/v2/pkg/db/dbconfig"
	"github.com/saucon/sauron/v2/pkg/env"
	"github.com/saucon/sauron/v2/pkg/env/envconfig"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

type Config struct {
	EnvConfig  envconfig.Config `mapstructure:"envLib"`
	AppEnvMode AppEnvMode       `mapstructure:"appEnvMode"`
	DBConfig   dbconfig.Config  `mapstructure:"databaseConfig"`
	LogConfig  logconfig.Config `mapstructure:"logConfig"`
	ConfigEnv  env.EnvConfig

	External struct {
	}

	Redis struct {
		Enable      bool   `mapstructure:"enable"`
		User        string `mapstructure:"user"`
		Password    string `mapstructure:"password"`
		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		MaxRetries  int    `mapstructure:"maxRetries"`
		DB          int    `mapstructure:"db"`
		UsingTLS    bool   `mapstructure:"usingTLS"`
		TLSCACert   string `mapstructure:"tlsCACert"`
		TLSCert     string `mapstructure:"tlsCert"`
		TLSKey      string `mapstructure:"tlsKey"`
		PoolSize    int    `mapstructure:"poolSize"`
		MinIdleConn int    `mapstructure:"minIdleConn"`
	}

	Request struct {
	}

	Server struct {
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"cleanup_period_seconds"`
			GracePeriodSeconds   int64 `mapstructure:"grace_period_seconds"`
		} `mapstructure:"shutdown"`
		Timeout struct {
			Duration time.Duration `mapstructure:"duration"`
		}
	} `mapstructure:"server"`

	Cors CORSConfig `mapstructure:"cors"`
}

type AppEnvMode struct {
	Mode           string `mapstructure:"mode"`        // dev, prod, dev_test
	GinMode        string `mapstructure:"ginMode"`     // debug, release
	DebugMode      bool   `mapstructure:"debugMode"`   // true, false
	IsPrettyLog    bool   `mapstructure:"isPrettyLog"` // true, false
	TestPathPrefix string `mapstructure:"testPathPrefix"`
}

type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allowOrigins"`
	AllowMethods     []string `mapstructure:"allowMethods"`
	AllowHeaders     []string `mapstructure:"allowHeaders"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
	ExposeHeaders    []string `mapstructure:"exposeHeaders"`
}
