package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/artem-xox/go-shorty/internal/server"
	"github.com/artem-xox/go-shorty/internal/service"
	"github.com/artem-xox/go-shorty/internal/store"
	"github.com/go-redis/redis/v8"
)

func main() {

	// serive config
	cfg, err := server.NewConfig()
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
		Store: &store.MemoryCacheStore{},
	}

	// test the Redis client connection by pinging the server
	pong, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis client connected: ", pong)

	http.HandleFunc("/ping", shortyService.Ping)

	http.HandleFunc("/getlink", shortyService.GetLink)
	http.HandleFunc("/setlink", shortyService.SetLink)

	// start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
