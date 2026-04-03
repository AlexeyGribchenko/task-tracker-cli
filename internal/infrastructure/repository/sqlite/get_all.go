package sqlite

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
)

const queryGetAll = `
	SELECT id, name, description, status, created_at, updated_at
	FROM tasks
`

func (s *Storage) GetTasks() ([]domain.Task, error) {
	const op = "repository.sqlite"

	rows, err := s.db.Query(queryGetAll)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var task domain.Task
		var statusStr string

		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&statusStr,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}

		task.Status = domain.TaskStatus(statusStr)

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration error: %w", op, err)
	}

	return tasks, nil
}
