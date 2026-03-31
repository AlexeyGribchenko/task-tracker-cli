package sqlite

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
)

const queryCreateTask = `
	INSERT INTO tasks (name, description, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id
`

func (s *Storage) CreateTask(task domain.Task) (*domain.Task, error) {

	err := s.db.QueryRow(queryCreateTask,
		task.Name,
		utils.ValueFromPointer(task.Description),
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&task.ID)

	if err != nil {
		return nil, fmt.Errorf("storage: failed to create task: %w", err)
	}

	return &task, nil
}
