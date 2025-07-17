package streaming

import (
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func NewKafkaWriter(kafkaBroker string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    "telemetry",
		Balancer: &kafka.Hash{}, // Uses IMEI as key for partitioning
	}
}

func GetKafkaWriter() *kafka.Writer {
	kafkaBroker := os.Getenv("KAFKA_BROKERS")
	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKERS environment variable is required")
	}

	if writer == nil {
		return NewKafkaWriter(kafkaBroker)
	}
	return writer
}
