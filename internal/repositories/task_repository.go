package repositories

import (
	"database/sql"
	"log"

	"github.com/gauravst/todo-backend-go/internal/models"
)

// TaskRepository defines the interface for user-related database operations
type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	GetAllTask() ([]*models.Task, error)
	UpdateTask(id int, task *models.Task) error
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
	query := `INSERT INTO todo (task) VALUES ($1) RETURNING *`
	row := r.db.QueryRow(query, task.Task)

	// Log the query and arguments
	log.Printf("Executing query: %s with task: %s", query, task.Task)

	// Scan the result
	err := row.Scan(
		&task.ID,
		&task.Task,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return err
	}

	// Log the populated task object
	log.Printf("Task after Scan: %+v", task)

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

func (r *taskRepository) GetAllTask() ([]*models.Task, error) {
	query := `SELECT id, task, status FROM todo`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(&task.ID, &task.Task, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTask(id int, task *models.Task) error {
	query := `UPDATE todo SET task = $1, status = $2 WHERE id = $3`
	row := r.db.QueryRow(query, task.Task, task.Status, task.ID)

	err := row.Scan(
		&task.ID,
		&task.Task,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
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
