package models

import (
	"time"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	RequestID int64     `json:"request_id"`
	Customer  string    `json:"customer"`
	Quantity  uint      `json:"quantity"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type ReqOrder struct {
	RequestID int64    `json:"request_id"`
	Data      []*Order `json:"data"`
}
