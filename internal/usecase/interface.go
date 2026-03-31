package usecase

import (
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock_interfaces.go -package=mocks

type TaskRepository interface {
	GetTasks() ([]domain.Task, error)
	CreateTask(task domain.Task) (*domain.Task, error)
	UpdateTaskStatus(id int, status domain.TaskStatus) error
}

type GetTasksUseCase interface {
	Execute() ([]domain.Task, error)
}

type CreateTaskUseCase interface {
	Execute(input dto.CreateTask) (*domain.Task, error)
}

type UpdateTaskStatusUseCase interface {
	Execute(input dto.UpdateTask) error
}
