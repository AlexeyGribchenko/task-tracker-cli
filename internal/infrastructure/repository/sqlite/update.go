package sqlite

import (
	"fmt"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
)

const queryUpdateTaskStatus = `
	UPDATE tasks
	SET status = $1, updated_at = $2
	WHERE id = $3
`

func (s *Storage) UpdateTaskStatus(id int, status domain.TaskStatus) error {

	result, err := s.db.Exec(queryUpdateTaskStatus,
		string(status),
		time.Now().UTC(),
		id,
	)

	if err != nil {
		return fmt.Errorf("storage: failed to update task status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("storage: failed get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("storage: task with id %d not found", id)
	}

	return nil
}
