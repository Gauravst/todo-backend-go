package repositories

import (
	"database/sql"

	"github.com/gauravst/todo-backend-go/internal/models"
)

// TaskRepository defines the interface for user-related database operations
type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db *sql.DB
}

// NewTaskRepository creates a new instance of userRepository
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO todo (task) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(query, task.Task).Scan(&task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetTaskByID(id int) (*models.Task, error) {
	task := &models.Task{}
	query := `SELECT id, task, status FROM todo WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Task, &task.Status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	query := `UPDATE todo SET task = $1, status = $2 WHERE id = $3`
	_, err := r.db.Exec(query, task.Task, task.Status, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(id int) error {
	query := `DELETE FROM todo WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
