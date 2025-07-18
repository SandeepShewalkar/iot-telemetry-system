package consumer

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"telemetry-ingestor/internal/model"
	"telemetry-ingestor/internal/storage"
	"telemetry-ingestor/internal/utils"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaBroker  string
	kafkaTopic   string
	kafkaGroupID string
)

const (
	keyPrefix string = "device:"
)

func StartKafkaConsumer(ctx context.Context) {

	validateAndSetKafkaEnv()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   kafkaTopic,
		GroupID: kafkaGroupID,
	})
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("Kafka consumer shutting down gracefully...")
				return
			}
			log.Println("Error reading message:", err)
			continue
		}

		var t model.Telemetry
		if err := json.Unmarshal(m.Value, &t); err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}

		processTelemetry(t, ctx)
	}
}

func processTelemetry(t model.Telemetry, ctx context.Context) {

	redisClient := storage.GetRedisClient()
	key := keyPrefix + t.IMEI
	lastData, _ := redisClient.HMGet(ctx, key, "lat", "lon", "dist").Result()

	var totalDist float64
	if lastData[0] != nil && lastData[1] != nil {
		prevLat, _ := strconv.ParseFloat(lastData[0].(string), 64)
		prevLon, _ := strconv.ParseFloat(lastData[1].(string), 64)
		dist := utils.Haversine(prevLat, prevLon, t.Latitude, t.Longitude)

		if lastData[2] != nil {
			totalDist, _ = strconv.ParseFloat(lastData[2].(string), 64)
		}
		totalDist += dist
	}

	redisClient.HSet(ctx, key, map[string]interface{}{
		"lat":  t.Latitude,
		"lon":  t.Longitude,
		"dist": totalDist,
	})

	log.Printf("Processed IMEI: %s → Total Distance: %.2f meters\n", t.IMEI, totalDist)
}

func validateAndSetKafkaEnv() {
	kafkaBroker = os.Getenv("KAFKA_BROKERS")
	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKERS environment variable is required")
	}
	kafkaTopic = os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		log.Fatal("KAFKA_BROKERS environment variable is required")
	}
	kafkaGroupID = os.Getenv("KAFKA_GROUPID")
	if kafkaTopic == "" {
		log.Fatal("KAFKA_GROUPID environment variable is required")
	}
}
