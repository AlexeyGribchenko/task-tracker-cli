package remove

import (
	"fmt"
	"strconv"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands"
)

//go:generate mockgen -source=remove.go -destination=mocks/remove_mock.go -package=mocks
type TaskRemover interface {
	Execute(input dto.RemoveTask) error
}

type SuccesWriter interface {
	PrintSuccessMessage(message string)
}

type CommandRemove struct {
	uc TaskRemover
	wr SuccesWriter
}

func New(uc TaskRemover, wr SuccesWriter) *CommandRemove {
	return &CommandRemove{
		uc: uc,
		wr: wr,
	}
}

func (c *CommandRemove) Execute(args []string) error {

	input, err := parseRemoveArgs(args)
	if err != nil {
		return fmt.Errorf("Failed to parse remove args: %w", err)
	}

	err = c.uc.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to delete task: %w", err)
	}

	c.wr.PrintSuccessMessage(fmt.Sprintf("Task succesfully removed: %d", input.ID))

	return nil
}

func parseRemoveArgs(args []string) (dto.RemoveTask, error) {

	if len(args) < 1 {
		return dto.RemoveTask{}, commands.ErrNotEnoughArguments
	}

	idStr := args[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.RemoveTask{}, domain.ErrInvalidTaskID
	}

	input := dto.RemoveTask{
		ID: id,
	}

	return input, nil
}
