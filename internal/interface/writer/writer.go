package writer

import (
	"os"
	"strings"

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
		// It just fixes bug with rendering on linux
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
		name = strings.ToLower(name)
		if isColumnNameValid(name) {
			headers = append(headers, name)
		}
	}
	headers = append(headers, ValidStatusName)

	table.Header(headers)

	return &TableWriter{table, headers}
}

func isColumnNameValid(name string) bool {
	switch name {
	case ValidIdName, ValidTaskNameName, ValidDescriptionName, ValidCreatedName, ValidUpdatedName, ValidStatusName:
		return true
	}
	return false
}

func (tw *TableWriter) AddRow(row []string) {
	tw.writer.Append(row)
}

func (tw *TableWriter) Render() error {
	return tw.writer.Render()
}

// func (cli *TableWriter) Print(content string) {
// 	fmt.Fprintln(cli.writer, content)
// }

// func (cli *TableWriter) Flush() error {
// 	return cli.writer.Flush()
// }
