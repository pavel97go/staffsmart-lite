package models

import "time"

type Slot struct {
	ID      int64     `json:"id"`
	Venue   string    `json:"venue"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
	Busy    bool      `json:"busy"`
}
