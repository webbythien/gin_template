package service

import (
	"admin-panel/v1/app/tasks"
	"admin-panel/v1/pkg/workers"
	"fmt"
)

func SendMessageTask(message string, room_id string) error {
	fmt.Println("zo service")
	return workers.Delay(
		"social_queue",
		"Worker.SendMessage",
		tasks.SaveMessageToRedis,
		message,
		room_id,
	)

	//return workers.Delay(
	//	"social_queue",
	//	"Worker.HealthCheck",
	//	tasks.HealthCheck,
	//	int64(8),
	//)
}
