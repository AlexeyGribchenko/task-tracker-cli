package cli

import (
	"errors"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/fatih/color"
)

var (
	ErrGetTasksFailed = errors.New("Failed to get list of tasks")
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

	// colorConfig := renderer.ColorizedConfig{
	// 	Header: renderer.Tint{
	// 		FG: renderer.Colors{color.Bold},
	// 		BG: renderer.Colors{color.ResetBlinking},
	// 	},
	// 	Column: renderer.Tint{
	// 		FG: renderer.Colors{color.FgHiWhite},
	// 	},
	// 	// It just fixes bug with rendering on linux
	// 	Border: renderer.Tint{
	// 		BG: renderer.Colors{color.ResetBlinking},
	// 	},
	// 	Separator: renderer.Tint{
	// 		BG: renderer.Colors{color.ResetBlinking},
	// 	},
	// }

	// table := tablewriter.NewTable(os.Stdout,
	// 	tablewriter.WithRenderer(renderer.NewColorized(colorConfig)),
	// 	tablewriter.WithRowMaxWidth(maxColumnWidth),
	// )
	// defer table.Render()

	// // TODO: make it configurable
	// table.Header([]string{"ID", "Task name", "description", "Created", "Updated", "Status"})

	for _, task := range tasks {
		status := task.Status

		statusStr := task.Status.String()

		switch status {
		case domain.TaskStatusCreated:
			statusStr = color.HiBlueString(statusStr)
		case domain.TaskStatusInProgress:
			statusStr = color.HiYellowString(statusStr)
		case domain.TaskStatusCompleted:
			statusStr = color.HiGreenString(statusStr)
		case domain.TaskStatusCancelled:
			statusStr = color.HiRedString(statusStr)
		}

		row := []string{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			utils.ValueFromPointer(task.Description),
			task.CreatedAt.Format("15:04 02.01"),
			task.UpdatedAt.Format("15:04 02.01"),
			statusStr,
		}

		a.writer.AddRow(row)
		// table.Append(row)
	}

	return a.writer.Render()
}
