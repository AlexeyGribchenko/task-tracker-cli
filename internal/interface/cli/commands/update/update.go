package update

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands"
)

//go:generate mockgen -source=update.go -destination=mocks/update_mock.go -package=mocks
type TaskUpdater interface {
	Execute(input dto.UpdateTask) error
}

type SuccesWriter interface {
	PrintSuccessMessage(message string)
}

type CommandUpdate struct {
	uc TaskUpdater
	wr SuccesWriter
}

var (
	ErrNotEnoughArguments = errors.New("Not enough arguments")
)

func New(uc TaskUpdater, wr SuccesWriter) *CommandUpdate {
	return &CommandUpdate{
		uc: uc,
		wr: wr,
	}
}

func (c *CommandUpdate) Execute(args []string) error {

	input, err := parseUpdateArgs(args)
	if err != nil {
		return fmt.Errorf("Failed to parse status args: %w", err)
	}

	err = c.uc.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to update task status: %w", err)
	}

	c.wr.PrintSuccessMessage(fmt.Sprintf("Task status successuly updated: %d", input.ID))

	return nil
}

func parseUpdateArgs(args []string) (dto.UpdateTask, error) {

	if len(args) < 2 {
		return dto.UpdateTask{}, commands.ErrNotEnoughArguments
	}

	idStr, status := args[0], args[1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.UpdateTask{}, err
	}

	input := dto.UpdateTask{
		ID:     id,
		Status: status,
	}

	return input, nil
}
