package models

import (
	"database/sql"
	"time"
)

type Metric struct {
	Name  string    `json:"name"`
	Value float64   `json:"value"`
	Time  time.Time `json:"time"`
}

func GetMetrics() []Metric {
	// Simulate retrieving metrics from a database
	return []Metric{{Name: "cpu_usage", Value: 50.2, Time: time.Now()}}
}
