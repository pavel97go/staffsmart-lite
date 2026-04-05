package models

import "time"

type Order struct {
	ID           int64     `json:"id"`
	SlotID       int64     `json:"slot_id"`
	CustomerName string    `json:"customer_name"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateOrderInput struct {
	SlotID       int64  `json:"slot_id"`
	CustomerName string `json:"customer_name"`
}
