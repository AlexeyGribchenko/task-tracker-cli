package cli

import (
	"errors"
	"os"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/config"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase"
	"github.com/fatih/color"
)

var (
	ErrInvalidCommand     = errors.New("invalid command")
	ErrNotEnoughArguments = errors.New("not enough arguments")
)

// TODO: COMMAND ID set -s (status) -d (description) -n (name)
// TODO: COMMAND list -s (name|status|created|updated), -f {status}
const (
	CommandCreateTask    = "add"
	CommandSetTaskStatus = "status"
	CommandGetTaskList   = "list"
	CommandRemoveTask    = "rm"
)

type Repository interface {
	usecase.TaskGetter
	usecase.TaskCreator
	usecase.TaskUpdater
	usecase.TaskRemover
	usecase.TaskSortedGetter
}

type App struct {
	createUC    usecase.CreateTaskUseCase
	getAllUC    usecase.GetTasksUseCase
	updateUC    usecase.UpdateTaskStatusUseCase
	removeUC    usecase.RemoveTaskUseCase
	getSortedUC usecase.GetTasksSortedUseCase
	writer      *writer.TableWriter
}

func New(db Repository, cfg *config.Config) *App {

	getUC := usecase.NewGetTasksUseCase(db)
	createUC := usecase.NewCreateTaskUseCase(db)
	updateUC := usecase.NewUpdateTaskUseCase(db)
	removeUC := usecase.NewRemoveTaskUseCase(db)

	writer := writer.New(cfg.Format)

	return &App{
		createUC: createUC,
		getAllUC: getUC,
		updateUC: updateUC,
		removeUC: removeUC,
		writer:   writer,
	}
}

func (a *App) Run() error {

	args := os.Args

	if len(args) < 2 {
		return ErrNotEnoughArguments
	}

	command := args[1]
	args = os.Args[2:]

	switch command {
	case CommandGetTaskList:
		return a.List(args)
	case CommandCreateTask:
		return a.Create(args)
	case CommandSetTaskStatus:
		return a.Set(args)
	case CommandRemoveTask:
		return a.Remove(args)
	default:
		return errors.New("Invalid command: " + color.New(color.Bold).Sprint(command))
	}
}
