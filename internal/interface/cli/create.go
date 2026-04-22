package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/fatih/color"
)

func (a *App) Create(args []string) error {

	if len(args) == 0 {
		return ErrNotEnoughArguments
	}

	newFlags := flag.NewFlagSet("add", flag.ContinueOnError)
	var description string

	newFlags.StringVar(&description, "d", "", "description of task or special notes")
	newFlags.Parse(args[1:])

	name := strings.Trim(args[0], "\"")

	input := dto.CreateTask{
		Name:        name,
		Description: nil,
	}

	if description != "" {
		input.Description = &description
	}

	task, err := a.createUC.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to create task: %w", err)
	}

	fmt.Println(color.GreenString("Task succsessfuly created: " + task.Name))

	return nil
}
