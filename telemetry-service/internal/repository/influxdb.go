package repository

import (
	"context"
	"fmt"

	influxdb "github.com/influxdata/influxdb-client-go/v2"

	"telemetry-service/internal/models"
)

type TelemetryRepository struct {
	client influxdb.Client
	org    string
	bucket string
}

func NewTelemetryRepository(url, token, org, bucket string) (*TelemetryRepository, error) {
	client := influxdb.NewClient(url, token)
	//resp, err := client.Setup(context.Background(), "admin", "password", org, bucket, 1)
	//if err != nil {
	//	return nil, err
	//}
	//log.Println("resp", resp)
	return &TelemetryRepository{client: client, org: org, bucket: bucket}, nil
}

func (r *TelemetryRepository) SaveTelemetry(telemetry models.Telemetry) error {
	writeAPI := r.client.WriteAPIBlocking(r.org, r.bucket)

	p := influxdb.NewPointWithMeasurement("telemetry").
		AddTag("device_id", telemetry.DeviceID).
		AddField("measure", telemetry.MeasureID).
		AddField("value", telemetry.MeasureValue).
		SetTime(telemetry.MeasuredAt)

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return fmt.Errorf("failed to write point: %w", err)
	}
	return nil
}

func (r *TelemetryRepository) GetTelemetry(deviceID string) ([]models.Telemetry, error) {
	queryAPI := r.client.QueryAPI("")
	query := `from(bucket:"your-bucket") |> range(start: -1h) |> filter(fn: (r) => r._measurement == "telemetry" and r.device_id == "` + deviceID + `")`

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var telemetries []models.Telemetry
	for result.Next() {
		if result.Err() != nil {
			return nil, result.Err()
		}

		telemetries = append(telemetries, models.Telemetry{
			DeviceID:     result.Record().ValueByKey("device_id").(string),
			MeasureID:    result.Record().ValueByKey("measure_id").(string),
			MeasureValue: result.Record().ValueByKey("measure_value").(string),
			MeasuredAt:   result.Record().Time(),
		})
	}

	return telemetries, nil
}
