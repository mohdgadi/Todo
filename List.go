package main

import "time"

// List Entity.
type List struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdat"`
	Tasks     []Task    `json:"tasks"`
}
