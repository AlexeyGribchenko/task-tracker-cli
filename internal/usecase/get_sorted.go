package usecase

import (
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

type GetTasksSortedUseCaseImpl struct {
	db TaskSortedGetter
}

func NewGetTasksSorted(db TaskSortedGetter) GetTasksSortedUseCase {
	return &GetTasksSortedUseCaseImpl{
		db: db,
	}
}

func (uc *GetTasksSortedUseCaseImpl) Execute(input dto.GetTasksSorted) ([]domain.Task, error) {
	const op = "usecase.GetTasksSorted"

	columnName, err := domain.ParseColumnName(input.ColumnSorted)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse column name: %w", op, err)
	}

	tasks, err := uc.db.GetSorted(columnName)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get sorted tasks: %w", op, err)
	}

	return tasks, nil
}
