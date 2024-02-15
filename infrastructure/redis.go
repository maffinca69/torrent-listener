package infrastructure

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var clientInstance *redis.Client

func RedisClient() *redis.Client {
	if clientInstance == nil {
		clientInstance = setupClient()
	}

	return clientInstance
}

func setupClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	index := os.Getenv("REDIS_DB_INDEX")
	dbIndex, err := strconv.Atoi(index)
	if err != nil {
		panic("Invalid database index")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       dbIndex,  // use default DB
	})

	return client
}
