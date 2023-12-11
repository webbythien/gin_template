package tasks

import (
	"admin-panel/v1/pkg/config"
	"admin-panel/v1/platform/cache"
	"fmt"
)

func SaveMessageToRedis(message string, room_id string) error {
	fmt.Println("zo task")

	emitter := cache.New(cache.Options{
		Host:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPass,
	})
	emitter.In(room_id).Emit("message", message)
	return nil
}
