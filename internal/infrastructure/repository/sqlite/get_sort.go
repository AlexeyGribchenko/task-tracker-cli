package sqlite

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
)

const queryGetSorted = `
	SELECT id, name, description, created_at, updated_at, status
	FROM tasks ORDER BY "$1"
`

func (s *Storage) GetSorted(columnName domain.ColumnName) ([]domain.Task, error) {

	rows, err := s.db.Query(queryGetSorted, string(columnName))
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted tasks: %w", err)
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
			&task.CreatedAt,
			&task.UpdatedAt,
			&statusStr,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to get sorted tasks: %w", err)
		}

		task.Status = domain.TaskStatus(statusStr)

		tasks = append(tasks, task)
	}

	return tasks, nil
}
