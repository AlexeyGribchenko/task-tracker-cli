package usecase

import "github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"

type RemoveTaskUseCaseImpl struct {
	db TaskRemover
}

func NewRemoveTaskUseCase(db TaskRemover) RemoveTaskUseCase {
	return &RemoveTaskUseCaseImpl{
		db: db,
	}
}

func (r *RemoveTaskUseCaseImpl) Execute(input dto.RemoveTask) error {

	id := input.ID
	if id < 0 {
		// А может в domain вынести и ошибки и валидацию id`шника`
		return ErrInvalidTaskID
	}

	return r.db.RemoveTask(id)
}
