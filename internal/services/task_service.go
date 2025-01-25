package services

import (
	"errors"
	"fmt"

	"github.com/gauravst/todo-backend-go/internal/models"
	"github.com/gauravst/todo-backend-go/internal/repositories"
)

type TaskService interface {
	CreateTask(task models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(task models.Task) error {
	if task.Task == "" {
		return errors.New("task cannot be empty")
	}

	err := s.taskRepo.CreateTask(&task)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *taskService) GetTaskByID(id int) (*models.Task, error) {
	user, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return user, nil
}

func (s *taskService) UpdateTask(task models.Task) error {
	if task.Task == "" {
		return errors.New("task cannot be empty")
	}

	err := s.taskRepo.UpdateTask(&task)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (s *taskService) DeleteTask(id int) error {
	err := s.taskRepo.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
