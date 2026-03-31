package usecase

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
)

type GetTaskUseCaseImpl struct {
	db TaskRepository
}

func NewGetTasksUseCase(db TaskRepository) *GetTaskUseCaseImpl {
	return &GetTaskUseCaseImpl{
		db: db,
	}
}

var _ GetTasksUseCase = (*GetTaskUseCaseImpl)(nil)

func (uc *GetTaskUseCaseImpl) Execute() ([]domain.Task, error) {

	tasks, err := uc.db.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("usecase.get_all: failed to get tasks from db: %w", err)
	}

	return tasks, nil
}
