package infra

import (
	"boilerplate/config"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"os"
)

type Redis struct {
	Client *redis.Client
}

// RedisNewClient create new instance of redis
func RedisNewClient(config *config.Config) *Redis {
	if !config.Redis.Enable {
		return nil
	}

	redisOpt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Username:     config.Redis.User,
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		MaxRetries:   config.Redis.MaxRetries,
		PoolFIFO:     false,
		PoolSize:     config.Redis.PoolSize,
		MinIdleConns: config.Redis.MinIdleConn,
	}

	if config.Redis.UsingTLS {
		caCert, err := os.ReadFile(config.Redis.TLSCACert)
		if err != nil {
			log.Fatal().Err(err).Msg("read file tls ca cert failed")
			return nil
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		cert, err := tls.LoadX509KeyPair(config.Redis.TLSCert, config.Redis.TLSKey)
		if err != nil {
			log.Fatal().Err(err).Msg("pair file tls cert and tls key failed")
			return nil
		}

		redisOpt.TLSConfig = &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		}
	}

	client := redis.NewClient(redisOpt)

	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		log.Error().Err(err).Msg("connection redis error")
		return nil
	}

	return &Redis{
		Client: client,
	}
}
