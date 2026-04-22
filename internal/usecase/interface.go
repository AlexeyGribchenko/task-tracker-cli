package usecase

import (
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

// -- Repository --
//
//go:generate mockgen -source=interface.go -destination=mocks/interfaces_mock.go -package=mocks
type TaskCreator interface {
	CreateTask(task domain.Task) (*domain.Task, error)
}

type TaskUpdater interface {
	UpdateTaskStatus(id int, status domain.TaskStatus) error
}

type TaskGetter interface {
	GetTasks() ([]domain.Task, error)
}

type TaskRemover interface {
	RemoveTask(id int) error
}

type TaskSortedGetter interface {
	GetSorted(columnName domain.ColumnTitle) ([]domain.Task, error)
}

type TaskFilteredGetter interface {
	GetFiltered(columnName domain.ColumnTitle, filterValue string) ([]domain.Task, error)
}

// -- UseCases ---

type GetTasksUseCase interface {
	Execute(input dto.GetTaskList) ([]domain.Task, error)
}

type CreateTaskUseCase interface {
	Execute(input dto.CreateTask) (*domain.Task, error)
}

type UpdateTaskStatusUseCase interface {
	Execute(input dto.UpdateTask) error
}

type RemoveTaskUseCase interface {
	Execute(input dto.RemoveTask) error
}
