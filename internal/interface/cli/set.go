package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/fatih/color"
)

var (
	ErrInvalidTaskID      = errors.New("Invalid task id!")
	ErrStatusUpdateFailed = errors.New("Failed to update task status!")
)

func (a *App) Set(args []string) error {

	idStr, status := args[0], args[1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ErrInvalidTaskID
	}

	input := dto.UpdateTask{
		ID:     id,
		Status: status,
	}

	err = a.updateUC.Execute(input)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrStatusUpdateFailed, err)
	}

	fmt.Println(color.GreenString("Task status successuly updated!"))

	return nil
}
