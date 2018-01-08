package main

import (
	"time"
)

// Task Entity.
type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdat"`
	Status    bool      `json:"status"`
}
