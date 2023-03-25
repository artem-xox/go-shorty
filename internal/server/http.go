package server

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"github.com/artem-xox/go-shorty/internal/service"
	"github.com/artem-xox/go-shorty/internal/store"
)

func StartHTTPServer() {

	logger := logrus.New()

	// serive config
	logger.Info("Config initiliazation...")
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("config init")
	}

	// create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DataBase,
	})

	shortyService := service.ShortyService{
		Store:  store.NewRedisCacheStore(rdb),
		Logger: logger,
	}

	// test the Redis client connection by pinging the server
	logger.Info("Redis connecting...")
	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ping", shortyService.Ping)
	http.HandleFunc("/set", shortyService.SetLink)
	http.HandleFunc("/l/", shortyService.GetLink)

	logger.Info("Listening and serving...")
	logger.Fatal(http.ListenAndServe(cfg.Service.Addr, nil))

}
