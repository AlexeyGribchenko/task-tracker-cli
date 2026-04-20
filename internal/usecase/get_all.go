package usecase

import (
	"fmt"
	"slices"
	"sort"

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

	if input.SortBy == "" {
		return tasks, nil
	}

	columnName, err := domain.ParseColumnName(input.SortBy)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse column name: %w", err)
	}

	switch columnName {
	case domain.ColumnName:
		sort.Slice(tasks, func(i, j int) bool {
			if input.Desc {
				return tasks[i].Name > tasks[j].Name
			}
			return tasks[i].Name < tasks[j].Name
		})
	case domain.ColumnStatus:
		sort.Slice(tasks, func(i, j int) bool {
			if input.Desc {
				return tasks[i].Status.String() > tasks[j].Status.String()
			}
			return tasks[i].Status.String() < tasks[j].Status.String()
		})
	case domain.ColumnCreatedAt:
		sort.Slice(tasks, func(i, j int) bool {
			if input.Desc {
				return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
			}
			return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
		})
	case domain.ColumnUpdatedAt:
		sort.Slice(tasks, func(i, j int) bool {
			if input.Desc {
				return tasks[i].UpdatedAt.Before(tasks[j].UpdatedAt)
			}
			return tasks[i].UpdatedAt.After(tasks[j].UpdatedAt)
		})
	case domain.ColumnDescription:
		sort.Slice(tasks, func(i, j int) bool {
			if input.Desc {
				return utils.ValueFromPointer(tasks[i].Description) > utils.ValueFromPointer(tasks[j].Description)
			}
			return utils.ValueFromPointer(tasks[i].Description) < utils.ValueFromPointer(tasks[j].Description)
		})
	}

	return tasks, nil
}

func (uc *GetTasksUseCaseImpl) filter(tasks []domain.Task, input dto.GetTaskList) ([]domain.Task, error) {

	if input.Status == "" {
		return tasks, nil
	}

	status, err := domain.ParseStatus(input.Status)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse status: %w", ErrInvalidTaskStatus)
	}

	tasks = slices.DeleteFunc(tasks, func(t domain.Task) bool {
		return t.Status != status
	})

	return tasks, nil
}
