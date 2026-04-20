package sqlite

import "fmt"

const queryRemoveById = `
	DELETE
	FROM tasks
	WHERE id = $1
`

func (s *Storage) RemoveTask(id int) error {

	res, err := s.db.Exec(queryRemoveById, id)

	if err != nil {
		return fmt.Errorf("storage: failed to delete task: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("storage: failed get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("storage: task with id %d not found", id)
	}

	return nil
}
