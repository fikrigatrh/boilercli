package config

import (
	"boilerplate/utils"
	"errors"
	"fmt"
	"github.com/saucon/sauron/v2/pkg/env"
	logDefault "log"
	"os"
)

func ProvideConfig() *Config {
	mode := os.Getenv("APP_MODE")
	cfg := new(Config)
	cfg.AppEnvMode.Mode = mode
	cfg.ConfigEnv = ProvideEnv(cfg)

	return cfg
}

func ProvideEnv(cfg *Config) (envCfg env.EnvConfig) {
	switch cfg.AppEnvMode.Mode {
	case utils.DEV:
		fmt.Printf("Run this app in : %s Environment \n", utils.DEV)
		err := os.Setenv("GIN_MODE", "debug")
		if err != nil {
			logDefault.Fatal(err.Error())
		}
		envCfg, err := env.New(utils.EnvDevFile, &cfg)
		if err != nil {
			logDefault.Fatal(err.Error())
		}
		return envCfg
	case utils.DEV_TEST:
		fmt.Printf("Run this app in : %s Environment \n", utils.DEV_TEST)
		err := os.Setenv("GIN_MODE", "debug")
		if err != nil {
			logDefault.Fatal(err.Error())
		}
		envCfg, err := env.New(cfg.AppEnvMode.TestPathPrefix+utils.EnvFile, &cfg)
		if err != nil {
			logDefault.Fatal(err)
		}
		return envCfg
	case utils.PROD:
		err := os.Setenv("GIN_MODE", "release")
		if err != nil {
			logDefault.Fatal(err.Error())
		}
		fmt.Printf("Run this app in : %s Environment \n", utils.PROD)
		envCfg, err := env.New(utils.EnvProdFile, &cfg)
		if err != nil {
			logDefault.Fatal(err)
		}
		return envCfg
	default:
		fmt.Printf("Run this app in Unknown Env \n")
		logDefault.Fatal(errors.New("Run this app in Unknown Env \n"))
	}

	return
}
