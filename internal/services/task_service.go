package services

import (
	"errors"
	"fmt"

	"github.com/gauravst/todo-backend-go/internal/models"
	"github.com/gauravst/todo-backend-go/internal/repositories"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	GetAllTask() ([]*models.Task, error)
	UpdateTask(id int, task *models.Task) error
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

func (s *taskService) CreateTask(task *models.Task) error {
	if task.Task == "" {
		return errors.New("task cannot be empty")
	}

	err := s.taskRepo.CreateTask(task)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *taskService) GetTaskByID(id int) (*models.Task, error) {
	task, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return task, nil
}

func (s *taskService) GetAllTask() ([]*models.Task, error) {
	tasks, err := s.taskRepo.GetAllTask()
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(id int, task *models.Task) error {
	if task.Task == "" {
		return errors.New("task cannot be empty")
	}

	err := s.taskRepo.UpdateTask(id, task)
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
