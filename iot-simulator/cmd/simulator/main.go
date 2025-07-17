package main

import (
	"iot-simulator/internal/simulator"
	"log"
	"os"
)

func main() {
	gatewayURL := os.Getenv("GATEWAY_URL")
	if gatewayURL == "" {
		log.Fatal("GATEWAY_URL environment variable is required")
	}
	simulator.StartSimulator(gatewayURL)
}
