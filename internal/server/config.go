package server

type RedisConfig struct {
	Addr string
}

type Config struct {
	Redis RedisConfig
}

func NewConfig() (*Config, error) {
	return &Config{
		RedisConfig{
			Addr: "0.0.0.0:6379",
		},
	}, nil
}
