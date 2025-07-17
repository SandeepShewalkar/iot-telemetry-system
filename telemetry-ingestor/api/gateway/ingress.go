package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"telemetry-ingestor/internal/model"
	"telemetry-ingestor/internal/streaming"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// TelemetryIngress godoc
// @Summary      Ingest telemetry data
// @Description  Accepts telemetry data in JSON format and publishes it to Kafka
// @Tags         Telemetry
// @Accept       json
// @Produce      json
// @Param        telemetry  body      model.Telemetry  true  "Telemetry Payload"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/telemetry [post]
func TelemetryIngress(c *gin.Context) {
	var t model.Telemetry
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, _ := json.Marshal(t)
	writer := streaming.GetKafkaWriter()
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(t.IMEI),
		Value: value,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka publish failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message queued"})
}

// TelemetryBatchIngress godoc
// @Summary     Ingest batch telemetry events
// @Description Accepts a batch of telemetry events from a single device (identified by IMEI)
// @Accept      json
// @Produce     json
// @Param       batch body model.TelemetryBatchRequest true "Batch telemetry payload"
// @Success     200 {object} map[string]interface{} "status and count of events queued"
// @Failure     400 {object} map[string]string "bad request error"
// @Failure     500 {object} map[string]string "internal server error"
// @Router      /v1/telemetrybatch [post]
func TelemetryBatchIngress(c *gin.Context) {
	var batch model.TelemetryBatchRequest
	if err := c.ShouldBindJSON(&batch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writer := streaming.GetKafkaWriter()
	messages := make([]kafka.Message, 0, len(batch.Events))

	for _, event := range batch.Events {
		event.IMEI = batch.IMEI // assign IMEI to each event
		value, err := json.Marshal(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal event"})
			return
		}

		messages = append(messages, kafka.Message{
			Key:   []byte(batch.IMEI),
			Value: value,
		})
	}

	if err := writer.WriteMessages(context.Background(), messages...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka publish failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "batch messages queued", "count": len(messages)})
}
