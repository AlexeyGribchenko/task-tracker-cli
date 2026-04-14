package cli

import (
	"errors"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
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

		row := make([]string, 0, 6)
		for _, field := range a.writer.HeaderFields {
			switch field {
			case writer.ValidIdName:
				row = append(row, fmt.Sprintf("%d", task.ID))
			case writer.ValidTaskNameName:
				row = append(row, task.Name)
			case writer.ValidDescriptionName:
				row = append(row, utils.ValueFromPointer(task.Description))
			case writer.ValidUpdatedName:
				row = append(row, task.UpdatedAt.Format("15:04 02.01"))
			case writer.ValidCreatedName:
				row = append(row, task.CreatedAt.Format("15:04 02.01"))
			case writer.ValidStatusName:
				row = append(row, statusStr)
			}
		}

		a.writer.AddRow(row)
	}

	return a.writer.Render()
}
