package cli

import (
	"errors"
	"os"

	// "github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase"
	"github.com/fatih/color"
)

var (
	ErrInvalidCommand     = errors.New("invalid command")
	ErrNotEnoughArguments = errors.New("not enough arguments")
)

// TODO: COMMAND set -s (status) -d (description) -n (name)
// TODO: COMMAND list sort -f (name|status|created|updated), -f {status} = ()
const (
	CommandCreateTask    = "add"
	CommandSetTaskStatus = "status"
	CommandGetTaskList   = "list"
)

type App struct {
	createUC usecase.CreateTaskUseCase
	getAllUC usecase.GetTasksUseCase
	updateUC usecase.UpdateTaskStatusUseCase
	writer   *writer.TableWriter
}

func New(
	cuc usecase.CreateTaskUseCase,
	guc usecase.GetTasksUseCase,
	uuc usecase.UpdateTaskStatusUseCase,
	wr *writer.TableWriter,
) *App {
	return &App{
		createUC: cuc,
		getAllUC: guc,
		updateUC: uuc,
		writer:   wr,
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
	default:
		return errors.New("Invalid command: " + color.New(color.Bold).Sprint(command))
	}
}
