package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/fatih/color"
)

var (
	ErrDeleteTask = errors.New("Failed to delete task!")
)

func (a *App) Remove(args []string) error {

	idStr := args[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ErrInvalidTaskID
	}

	input := dto.RemoveTask{ID: id}

	err = a.removeUC.Execute(input)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDeleteTask, err)
	}

	fmt.Println(color.GreenString("Task succesfully removed!"))

	return nil
}
