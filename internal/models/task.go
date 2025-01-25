package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Task      string    `json:"task" validate:"required"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
