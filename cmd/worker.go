package cmd

import (
	"admin-panel/v1/app/tasks"
	"admin-panel/v1/pkg/config"
	"admin-panel/v1/pkg/workers"
)

// Setting configures and returns a WorkerConfig for background workers.
func Setting(broker, resultBackend string) *config.WorkerConfig {
	return &config.WorkerConfig{
		Broker:        broker,        // Specify the broker URL, e.g., RabbitMQ or Redis.
		ResultBackend: resultBackend, // Specify the result backend URL, e.g., Redis or a database.
		Workers: map[string]config.Task{
			"social_queue": map[string]interface{}{
				"Worker.SendMessage": tasks.SaveMessageToRedis, // Configure a worker named "Worker.SendMessage" and assign it the SaveMessageToRedis function.
				"Worker.HealthCheck": tasks.HealthCheck,
			},
		},
	}
}

// WorkerExecute configures and executes background workers.
func WorkerExecute(queueName, consume string, concurrency int) error {
	// Configure the WorkerConfig using the Setting function.
	wcf := Setting(config.BrokerUrl, config.ResultBackend)

	// Create a WorkerConfig instance using the NewWorker function and provide the queueName and configured WorkerConfig.
	cnf := config.NewWorker(queueName, wcf)

	// Execute background workers with the specified queueName, consume, and concurrency.

	return workers.Execute(cnf, consume, concurrency)
}
