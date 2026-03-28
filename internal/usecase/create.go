package usecase

import (
	"errors"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

var (
	ErrEmptyTaskName = errors.New("empty task name!")
)

type CreateTaskUseCaseImpl struct {
	db TaskRepository
}

func NewCreateTaskUse(db TaskRepository) *CreateTaskUseCaseImpl {
	return &CreateTaskUseCaseImpl{
		db: db,
	}
}

var _ CreateTaskUseCase = (*CreateTaskUseCaseImpl)(nil)

func (uc *CreateTaskUseCaseImpl) Execute(input dto.CreateTask) (*domain.Task, error) {

	if input.Name == "" {
		return nil, ErrEmptyTaskName
	}

	task := domain.NewTask(input.Name, input.Description)

	task, err := uc.db.CreateTask(*task)
	if err != nil {
		return nil, fmt.Errorf("usecase.create: failed to create task in db: %w", err)
	}

	return task, nil
}
