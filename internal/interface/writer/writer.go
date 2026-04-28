package writer

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

const (
	ValidIdName          = "id"
	ValidTaskNameName    = "name"
	ValidDescriptionName = "description"
	ValidCreatedName     = "created"
	ValidUpdatedName     = "updated"
	ValidStatusName      = "status"
)

type TableWriter struct {
	writer       *tablewriter.Table
	HeaderFields []string
}

func New(cfg Config) *TableWriter {

	colorConfig := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.Bold},
			BG: renderer.Colors{color.ResetBlinking},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgHiWhite},
		},
		// FIXME: Workaround for Linux terminal rendering issue with table borders
		// Setting ResetBlinking prevents flickering/artifacts in some terminals
		Border: renderer.Tint{
			BG: renderer.Colors{color.ResetBlinking},
		},
		Separator: renderer.Tint{
			BG: renderer.Colors{color.ResetBlinking},
		},
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorConfig)),
		tablewriter.WithRowMaxWidth(cfg.MaxColumnWidth),
	)

	headers := []string{ValidIdName, ValidTaskNameName}
	for _, name := range cfg.ExtraColumns {
		if isColumnNameValid(name) {
			headers = append(headers, name)
		}
	}
	headers = append(headers, ValidStatusName)

	table.Header(headers)

	return &TableWriter{table, headers}
}

func isColumnNameValid(name string) bool {

	name = strings.ToLower(name)

	switch name {
	case ValidIdName, ValidTaskNameName, ValidDescriptionName, ValidCreatedName, ValidUpdatedName, ValidStatusName:
		return true
	}
	return false
}

func (tw *TableWriter) PrintSuccessMessage(message string) {
	fmt.Println(color.GreenString(message))
}

func (tw *TableWriter) RenderTable(tasks []domain.Task) error {
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
		case domain.TaskStatusActive:
			statusStr = color.HiYellowString(statusStr)
		case domain.TaskStatusCompleted:
			statusStr = color.HiGreenString(statusStr)
		case domain.TaskStatusCancelled:
			statusStr = color.HiRedString(statusStr)
		}

		row := make([]string, 0, 6)
		for _, header := range tw.HeaderFields {
			switch header {
			case ValidIdName:
				row = append(row, fmt.Sprintf("%d", task.ID))
			case ValidTaskNameName:
				row = append(row, task.Name)
			case ValidDescriptionName:
				row = append(row, utils.ValueFromPointer(task.Description))
			case ValidUpdatedName:
				row = append(row, task.UpdatedAt.Format("15:04 02.01"))
			case ValidCreatedName:
				row = append(row, task.CreatedAt.Format("15:04 02.01"))
			case ValidStatusName:
				row = append(row, statusStr)
			}
		}

		tw.writer.Append(row)
	}
	return tw.writer.Render()
}
