package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

var (
	ErrGetTasksFailed = errors.New("Failed to get list of tasks")
)

const (
	maxColumnWidth = 50
)

func (a *App) List(args []string) error {

	tasks, err := a.getAllUC.Execute()
	if err != nil {
		return ErrGetTasksFailed
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet...")
		return nil
	}

	colorConfig := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.Bold},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgWhite},
		},
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorConfig)),
		tablewriter.WithRowMaxWidth(maxColumnWidth),
	)
	// TODO: make it configurable
	table.Header("ID", "Task name", "Description", "Created", "Updated", "Status")

	for _, task := range tasks {
		status := task.Status

		statusStr := task.Status.String()

		switch status {
		case domain.TaskStatusCreated:
			statusStr = color.BlueString(statusStr)
		case domain.TaskStatusInProgress:
			statusStr = color.YellowString(statusStr)
		case domain.TaskStatusCompleted:
			statusStr = color.GreenString(statusStr)
		case domain.TaskStatusCancelled:
			statusStr = color.RedString(statusStr)
		}

		row := []string{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			utils.ValueFromPointer(task.Description),
			task.CreatedAt.Format("15:04 02.01"),
			task.UpdatedAt.Format("15:04 02.01"),
			statusStr,
		}
		table.Append(row)
	}
	table.Render()

	return nil
}
