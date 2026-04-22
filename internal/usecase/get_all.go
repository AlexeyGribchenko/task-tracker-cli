package usecase

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
)

type GetTasksUseCaseImpl struct {
	db TaskGetter
}

func NewGetTasksUseCase(db TaskGetter) GetTasksUseCase {
	return &GetTasksUseCaseImpl{
		db: db,
	}
}

func (uc *GetTasksUseCaseImpl) Execute(input dto.GetTaskList) ([]domain.Task, error) {

	tasks, err := uc.db.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("Failed to get tasks from db: %w", err)
	}

	tasks, err = uc.sort(tasks, input)
	if err != nil {
		return nil, fmt.Errorf("Failed to sort tasks: %w", err)
	}

	tasks, err = uc.filter(tasks, input)
	if err != nil {
		return nil, fmt.Errorf("Failed to filter tasks: %w", err)
	}

	return tasks, nil
}

func (uc *GetTasksUseCaseImpl) sort(tasks []domain.Task, input dto.GetTaskList) ([]domain.Task, error) {

	if input.SortBy == "" || input.SortBy == domain.ColumnId {
		return tasks, nil
	}

	columnName, err := domain.ParseColumnName(input.SortBy)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse column name: %w", err)
	}

	slices.SortFunc(tasks, func(a, b domain.Task) int {
		var result int

		switch columnName {
		case domain.ColumnName:
			result = cmp.Compare(a.Name, b.Name)
		case domain.ColumnCreatedAt:
			result = cmp.Compare(a.CreatedAt.Unix(), b.CreatedAt.Unix())
		case domain.ColumnUpdatedAt:
			result = cmp.Compare(a.UpdatedAt.Unix(), b.UpdatedAt.Unix())
		case domain.ColumnStatus:
			result = cmp.Compare(a.Status.String(), b.Status.String())
		case domain.ColumnDescription:
			result = cmp.Compare(
				utils.ValueFromPointer(a.Description),
				utils.ValueFromPointer(b.Description),
			)
		}

		if input.Desc {
			result = -result
		}

		return result
	})

	return tasks, nil
}

func (uc *GetTasksUseCaseImpl) filter(tasks []domain.Task, input dto.GetTaskList) ([]domain.Task, error) {

	if input.Status == "" {
		return tasks, nil
	}

	status, err := domain.ParseStatus(input.Status)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse status: %w", err)
	}

	tasks = slices.DeleteFunc(tasks, func(t domain.Task) bool {
		return t.Status != status
	})

	return tasks, nil
}
