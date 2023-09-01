package base

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var (
	DataStore *redis.Client
)

func InitKeyDB() {
	DataStore = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})

	err := DataStore.Ping().Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")
}

func InitRedis() {
	DataStore = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
