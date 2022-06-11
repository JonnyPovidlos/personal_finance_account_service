package models

import "time"

type Operation struct {
	UserId      int
	Type        TypeOperation
	Amount      float64
	Description string
	CategoryId  int
	CreatedAt   time.Time
}

type TypeOperation string

const (
	SALE TypeOperation = "sale"
	BUY                = "buy"
)
