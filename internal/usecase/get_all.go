package usecase

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
)

type GetTasksUseCaseImpl struct {
	db TaskGetter
}

func NewGetTasksUseCase(db TaskGetter) GetTasksUseCase {
	return &GetTasksUseCaseImpl{
		db: db,
	}
}

func (uc *GetTasksUseCaseImpl) Execute() ([]domain.Task, error) {

	tasks, err := uc.db.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("usecase.get_all: failed to get tasks from db: %w", err)
	}

	return tasks, nil
}
