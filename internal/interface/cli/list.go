package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/fatih/color"
)

var (
	ErrGetTasksFailed = errors.New("Failed to get list of tasks")
)

func (a *App) List(args []string) error {

	newFlags := flag.NewFlagSet("list", flag.ContinueOnError)

	var sortedColumn string
	var status string
	var desc bool

	newFlags.StringVar(&sortedColumn, "s", "", "name of a column that will be sorted")
	newFlags.StringVar(&status, "f", "", "status wich is used to filter task list")
	newFlags.BoolVar(&desc, "desc", false, "sorting order")
	newFlags.Parse(args)

	var tasks []domain.Task
	var err error

	input := dto.GetTaskList{
		Status: status,
		SortBy: sortedColumn,
		Desc:   desc,
	}

	tasks, err = a.getAllUC.Execute(input)

	if err != nil {
		return fmt.Errorf("%w: %w", ErrGetTasksFailed, err)
	}

	return a.render(tasks)
}

func (a *App) render(tasks []domain.Task) error {

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
