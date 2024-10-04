package models

import "time"

type Telemetry struct {
	DeviceID     string    `json:"device_id"`
	MeasureID    string    `json:"measure_id"`
	MeasureValue string    `json:"measure_value"`
	MeasuredAt   time.Time `json:"measured_at"`
}
