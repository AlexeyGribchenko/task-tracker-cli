package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/create"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/list"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/remove"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/update"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase"
)

// TODO: COMMAND ID set -s (status) -d (description) -n (name)
const (
	CommandCreateTask    = "add"
	CommandSetTaskStatus = "status"
	CommandGetTaskList   = "list"
	CommandRemoveTask    = "rm"
)

var (
	ErrNotEnoughArguments = errors.New("Not enough arguments")
)

type Repository interface {
	usecase.TaskGetter
	usecase.TaskCreator
	usecase.TaskUpdater
	usecase.TaskRemover
}

type Command interface {
	Execute(args []string) error
}

type App struct {
	commands map[string]Command
}

func New(db Repository, wr writer.TableWriter) *App {

	newCommands := map[string]Command{
		CommandCreateTask:    initCreateCommand(db, wr),
		CommandGetTaskList:   initListCommand(db, wr),
		CommandSetTaskStatus: initStatusCommand(db, wr),
		CommandRemoveTask:    initRemoveCommand(db, wr),
	}

	return &App{
		commands: newCommands,
	}
}

func (a *App) Run() error {

	args := os.Args

	if len(args) < 2 {
		return ErrNotEnoughArguments
	}

	commandName := args[1]

	cmd, commandFound := a.commands[commandName]
	if !commandFound {
		return fmt.Errorf("Invalid command: %s", commandName)
	}

	commandArgs := os.Args[2:]

	return cmd.Execute(commandArgs)
}

func initCreateCommand(db Repository, wr writer.TableWriter) *create.CommandCreate {
	return create.New(
		usecase.NewCreateTaskUseCase(db),
		&wr,
	)
}

func initListCommand(db Repository, wr writer.TableWriter) *list.CommandList {
	return list.New(
		usecase.NewGetTasksUseCase(db),
		&wr,
	)
}

func initRemoveCommand(db Repository, wr writer.TableWriter) *remove.CommandRemove {
	return remove.New(
		usecase.NewRemoveTaskUseCase(db),
		&wr,
	)
}

func initStatusCommand(db Repository, wr writer.TableWriter) *update.CommandUpdate {
	return update.New(
		usecase.NewUpdateTaskUseCase(db),
		&wr,
	)
}
