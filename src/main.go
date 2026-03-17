package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"pulsar-analytics-engine/src/controllers"
	"pulsar-analytics-engine/src/models"
	"pulsar-analytics-engine/src/utils"
)

func main() {
	http.HandleFunc("/metrics", controllers.GetMetrics)

	http.ListenAndServe(":8080", nil)
}
