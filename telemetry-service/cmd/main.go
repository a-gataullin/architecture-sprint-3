package main

import (
	"log"
	"net/http"
	"os"

	"telemetry-service/internal/handlers"
	"telemetry-service/internal/repository"
)

func main() {
	repo, err := repository.NewTelemetryRepository(os.Getenv("INFLUXDB_URL"), "UNUQ3AZGMC8QN6J31gTT_PNJ3Ts5sQbV9uBxykq371OcXGOyJ9csTNtWQmvsgZK4mdYGpZaf4UP84KNrCdbDTQ==", os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))
	if err != nil {
		panic(err)
	}
	telemetryHandler := handlers.NewTelemetryHandler(repo)

	http.HandleFunc("/telemetry", telemetryHandler.PostTelemetry)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
