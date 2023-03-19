package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {
	// create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// test the Redis client connection by pinging the server
	pong, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis client connected: ", pong)

	// define the HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// retrieve a value from Redis
		val, err := rdb.Get(rdb.Context(), "key").Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Value from Redis: %s", val)
	})

	// start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
