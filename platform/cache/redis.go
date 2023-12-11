package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.

var client *redis.ClusterClient

func init() {
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			os.Getenv("REDIS_NODE_1"),
			os.Getenv("REDIS_NODE_2"),
			os.Getenv("REDIS_NODE_3"),
		},
		Password: "",
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	log.Println("Connected to Redis Cluster:", pong)
	log.Println("[REDIS] Cluster connection successful")
}
