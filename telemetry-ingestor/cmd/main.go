package main

import (
	"telemetry-ingestor/api"
	_ "telemetry-ingestor/docs" // swag docs
	"telemetry-ingestor/internal/consumer"
)

func main() {

	go consumer.StartKafkaConsumer()

	api.StartAPIServer()
}
