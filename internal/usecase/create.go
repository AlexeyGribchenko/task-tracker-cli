package usecase

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

type CreateTaskUseCaseImpl struct {
	db TaskRepository
}

func NewCreateTaskUseCase(db TaskRepository) *CreateTaskUseCaseImpl {
	return &CreateTaskUseCaseImpl{
		db: db,
	}
}

var _ CreateTaskUseCase = (*CreateTaskUseCaseImpl)(nil)

func (uc *CreateTaskUseCaseImpl) Execute(input dto.CreateTask) (*domain.Task, error) {
	const op = "usecase.create.Execute"

	task, err := domain.NewTask(input.Name, input.Description)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	task, err = uc.db.CreateTask(*task)
	if err != nil {
		return nil, fmt.Errorf("usecase.create: failed to create task in db: %w", err)
	}

	return task, nil
}
