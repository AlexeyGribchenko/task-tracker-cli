package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

var (
	ErrTaskNameNotProvided = errors.New("Task name is not provided !")
	ErrInvalidTaskName     = errors.New("Task name should not be empty!")
	ErrCreateFailed        = errors.New("Failed to create task!")
)

func (a *App) Create(args []string) error {

	if len(args) == 0 {
		return ErrTaskNameNotProvided
	}

	newFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	var description string

	newFlags.StringVar(&description, "d", "", "description of task or special notes")
	newFlags.Parse(args[1:])

	name := strings.Trim(args[0], "\"")
	if name == "" {
		return ErrInvalidTaskName
	}

	input := dto.CreateTask{
		Name: name,
	}

	if description == "" {
		input.Description = nil
	} else {
		input.Description = &description
	}

	task, err := a.createUC.Execute(input)
	if err != nil {
		return ErrCreateFailed
	}

	fmt.Println("Task succsessfuly created: " + task.Name)

	return nil
}
