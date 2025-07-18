package main

import (
	"context"
	"telemetry-ingestor/api"
	_ "telemetry-ingestor/docs" // swag docs
	"telemetry-ingestor/internal/consumer"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumerCount := 3
	for i := 0; i < consumerCount; i++ {
		go consumer.StartKafkaConsumer(ctx)
	}
	api.StartAPIServer()
}
