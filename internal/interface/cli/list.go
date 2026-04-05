package cli

import (
	"errors"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
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

	// ID, Name, Description, Creation time, Update time, Status
	pattern := "%v\t%s\t%s\t%s\t%s\t%s"

	str := fmt.Sprintf(pattern, "ID", "Task name", "Description", "Created", "Updated", "Status")

	defer a.writer.Flush()
	a.writer.Print(str)

	for _, task := range tasks {
		a.writer.Print(fmt.Sprintf(pattern,
			task.ID,
			task.Name,
			utils.ValueFromPointer(task.Description),
			task.CreatedAt.Format("15:04 02.01"),
			task.UpdatedAt.Format("15:04 02.01"),
			task.Status,
		))
	}

	return nil
}
