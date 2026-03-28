package usecase

import (
	"errors"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

var (
	ErrTaskNotFound      = errors.New("task not found")
	ErrInvalidTaskStatus = errors.New("invalid task status")
	ErrInvalidTaskID     = errors.New("invalid task id")
)

type UpdateTaskStatusUseCaseImpl struct {
	db TaskRepository
}

func NewUpdateTaskUseCase(db TaskRepository) *UpdateTaskStatusUseCaseImpl {
	return &UpdateTaskStatusUseCaseImpl{
		db: db,
	}
}

var _ UpdateTaskStatusUseCase = (*UpdateTaskStatusUseCaseImpl)(nil)

func (uc *UpdateTaskStatusUseCaseImpl) Execute(input dto.UpdateTask) error {

	if input.ID <= 0 {
		return ErrInvalidTaskID
	}

	status, err := domain.ParseStatus(input.Status)
	if err != nil {
		return ErrInvalidTaskStatus
	}

	err = uc.db.UpdateTaskStatus(input.ID, status)
	if err != nil {
		return fmt.Errorf("usecase.update: failed to update task in db: %w", err)
	}

	return nil
}
