package config

import (
	"github.com/RichardKnop/machinery/v1/config"
)

// Worker represents a worker configuration.
type Worker struct {
	Task   map[string]interface{} // Task configuration for the worker.
	Config *config.Config         // Machinery worker configuration.
}

// Task represents a map of worker tasks.
type Task map[string]interface{}

// WorkerConfig represents the configuration for background workers.
type WorkerConfig struct {
	Broker        string          // Broker URL, e.g., RabbitMQ or other messaging systems.
	ResultBackend string          // Result backend URL, e.g., Redis or a database.
	Workers       map[string]Task // Map of worker tasks by queue name.
}

// NewWorker creates a new worker configuration.
func NewWorker(queueName string, wcf *WorkerConfig) *Worker {
	return &Worker{
		Config: &config.Config{
			Broker:          wcf.Broker,        // Set the broker URL.
			DefaultQueue:    queueName,         // Set the default queue name.
			ResultBackend:   wcf.ResultBackend, // Set the result backend URL.
			ResultsExpireIn: 3600,              // Set result expiration time in seconds.
			AMQP: &config.AMQPConfig{
				Exchange:      "machinery_exchange", // Set the exchange name.
				ExchangeType:  "direct",             // Set the exchange type.
				BindingKey:    queueName,            // Set the binding key (queue name).
				PrefetchCount: 3,                    // Set the prefetch count for concurrency.
			},
		},
		Task: wcf.Workers[queueName], // Get the task configuration for the specified queueName.
	}
}
