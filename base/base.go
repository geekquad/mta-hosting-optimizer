package base

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	DataStore *redis.Client
)

func InitKeyDB() {
	DataStore = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("KEYDB_HOST"), os.Getenv("KEYDB_PORT")),
		Password: "",
		DB:       0,
	})

	err := DataStore.Ping().Err()
	if err != nil {
		log.Fatalf("Failed to connect to KeyDB: %v", err)
	}
	fmt.Println("Connected to KeyDB!")
	InitDataStore()

}

func InitRedis() {
	DataStore = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
	err := DataStore.Ping().Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")
	InitDataStore()
}

func InitDataStore() {
	DataStore.HSet("mta-prod-1", "127.0.0.1", "true")
	DataStore.HSet("mta-prod-1", "127.0.0.2", "false")
	DataStore.HSet("mta-prod-2", "127.0.0.3", "true")
	DataStore.HSet("mta-prod-2", "127.0.0.4", "true")
	DataStore.HSet("mta-prod-2", "127.0.0.5", "false")
	DataStore.HSet("mta-prod-3", "127.0.0.6", "false")
}
