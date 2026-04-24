package remove

import (
	"fmt"
	"strconv"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

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

	err = c.uc.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to delete task: %w", err)
	}

	c.wr.PrintSuccessMessage("Task succesfully removed!")

	return nil
}

func parseRemoveArgs(args []string) (dto.RemoveTask, error) {

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
