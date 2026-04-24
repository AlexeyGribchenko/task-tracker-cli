package create

import (
	"flag"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands"
)

//go:generate mockgen -source=create.go -destination=mocks/create_mock.go -package=mocks
type TaskCreator interface {
	Execute(input dto.CreateTask) (*domain.Task, error)
}

type SuccesWriter interface {
	PrintSuccessMessage(message string)
}

type CommandCreate struct {
	uc TaskCreator
	wr SuccesWriter
}

func New(uc TaskCreator, wr SuccesWriter) *CommandCreate {
	return &CommandCreate{
		uc: uc,
		wr: wr,
	}
}

func (c *CommandCreate) Execute(args []string) error {

	if len(args) == 0 {
		return commands.ErrNotEnoughArguments
	}

	input, err := parseCreateArgs(args)
	if err != nil {
		return fmt.Errorf("Failed to parse cli arguments: %w", err)
	}

	task, err := c.uc.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to create task: %w", err)
	}

	c.wr.PrintSuccessMessage(fmt.Sprintf("Task succsessfuly created: %s", task.Name))

	return nil
}

func parseCreateArgs(args []string) (dto.CreateTask, error) {

	newFlags := flag.NewFlagSet("add", flag.ContinueOnError)
	var description string

	newFlags.StringVar(&description, "d", "", "description of task or special notes")

	err := newFlags.Parse(args[1:])
	if err != nil {
		return dto.CreateTask{}, fmt.Errorf("Failed to parse args: %w", err)
	}

	input := dto.CreateTask{
		Name:        args[0],
		Description: nil,
	}

	if description != "" {
		input.Description = &description
	}

	return input, nil
}
