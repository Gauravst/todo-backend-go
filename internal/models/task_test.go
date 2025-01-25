package models_test

import (
	"testing"
	"time"

	"github.com/gauravst/todo-backend-go/internal/models"
	"github.com/go-playground/validator/v10"
)

func TestUserValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		user    models.Task
		wantErr bool
	}{
		{
			name: "valid task",
			user: models.Task{
				Task:      "Setup a login system",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "missing task",
			user: models.Task{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("task validation failed: %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskTimestamps(t *testing.T) {
	user := models.Task{
		Task: "this is my task",
	}

	if !user.CreatedAt.IsZero() || !user.UpdatedAt.IsZero() {
		t.Error("expected CreatedAt and UpdatedAt to be zero values initially")
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if user.CreatedAt.IsZero() || user.UpdatedAt.IsZero() {
		t.Error("expected CreatedAt and UpdatedAt to be set")
	}

	if user.CreatedAt != now || user.UpdatedAt != now {
		t.Error("expected CreatedAt and UpdatedAt to match the current time")
	}
}
