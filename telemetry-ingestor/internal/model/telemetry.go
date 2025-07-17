package model

type Telemetry struct {
	IMEI       string  `json:"imei"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DeviceTime int64   `json:"device_time"`
}

type TelemetryBatchRequest struct {
	IMEI   string      `json:"imei"`
	Events []Telemetry `json:"events"`
}
