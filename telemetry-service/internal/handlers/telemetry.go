package handlers

import (
	"encoding/json"
	"net/http"

	"telemetry-service/internal/models"
	"telemetry-service/internal/repository"
)

type TelemetryHandler struct {
	repo *repository.TelemetryRepository
}

func NewTelemetryHandler(repo *repository.TelemetryRepository) *TelemetryHandler {
	return &TelemetryHandler{repo: repo}
}

func (h *TelemetryHandler) PostTelemetry(w http.ResponseWriter, r *http.Request) {
	var telemetry models.Telemetry

	if err := json.NewDecoder(r.Body).Decode(&telemetry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.SaveTelemetry(telemetry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TelemetryHandler) GetTelemetry(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("device_id")
	if deviceID == "" {
		http.Error(w, "device_id is required", http.StatusBadRequest)
		return
	}

	telemetryData, err := h.repo.GetTelemetry(deviceID)
	if err != nil {
		http.Error(w, "Error retrieving telemetry data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(telemetryData)
}
