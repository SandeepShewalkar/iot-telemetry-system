package simulator

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Telemetry struct {
	IMEI       string  `json:"imei"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DeviceTime int64   `json:"device_time"`
}

func StartSimulator(gatewayURL string) {
	imeis := []string{"123456789000001", "123456789000002", "123456789000003"}
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		for _, imei := range imeis {
			sendTelemetry(imei, gatewayURL)
		}
	}
}

func sendTelemetry(imei, gatewayURL string) {
	payload := Telemetry{
		IMEI:       imei,
		Latitude:   19.0 + rand.Float64(),
		Longitude:  72.0 + rand.Float64(),
		DeviceTime: time.Now().UnixMilli(),
	}

	data, _ := json.Marshal(payload)
	resp, err := http.Post(gatewayURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Failed to send telemetry:", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Sent telemetry for IMEI %s → Status: %s\n", imei, resp.Status)
}
