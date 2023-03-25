package server

import "os"

type ServiceConfig struct {
	Addr string
}

type RedisConfig struct {
	Addr     string
	Password string
	DataBase int
}

type Config struct {
	Service ServiceConfig
	Redis   RedisConfig
}

func NewConfig() (*Config, error) {
	return &Config{
		ServiceConfig{
			Addr: os.Getenv("SERVICE_ADDR"),
		},
		RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),
			DataBase: 1,
		},
	}, nil
}
