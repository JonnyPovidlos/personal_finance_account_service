package models

import "time"

type Operation struct {
	OperationType string
	Amount        float64
	Description   string
	Category      string
	CreatedAt     time.Time
}
