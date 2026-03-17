package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"pulsar-analytics-engine/src/models"
	"pulsar-analytics-engine/src/utils"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := models.GetMetrics()
	json.NewEncoder(w).Encode(metrics)
}
