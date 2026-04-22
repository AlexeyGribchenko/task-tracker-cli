package usecase

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

type CreateTaskUseCaseImpl struct {
	db TaskCreator
}

func NewCreateTaskUseCase(db TaskCreator) CreateTaskUseCase {
	return &CreateTaskUseCaseImpl{
		db: db,
	}
}

func (uc *CreateTaskUseCaseImpl) Execute(input dto.CreateTask) (*domain.Task, error) {

	task, err := domain.NewTask(input.Name, input.Description)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new task: %w", err)
	}

	task, err = uc.db.CreateTask(*task)
	if err != nil {
		return nil, fmt.Errorf("Failed to save task in database: %w", err)
	}

	return task, nil
}
