package list

import (
	"flag"
	"fmt"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
)

//go:generate mockgen -source=list.go -destination=mocks/list_mock.go -package=mocks
type TaskWriter interface {
	RenderTable(tasks []domain.Task) error
}

type TasksGetter interface {
	Execute(input dto.GetTaskList) ([]domain.Task, error)
}

type CommandList struct {
	uc TasksGetter
	wr TaskWriter
}

func New(uc TasksGetter, wr TaskWriter) *CommandList {
	return &CommandList{
		uc: uc,
		wr: wr,
	}
}

func (c *CommandList) Execute(args []string) error {

	input, err := parseListFlags(args)
	if err != nil {
		return fmt.Errorf("Failed to parse list flags: %w", err)
	}

	tasks, err := c.uc.Execute(input)
	if err != nil {
		return fmt.Errorf("Failed to get tasks: %w", err)
	}

	return c.wr.RenderTable(tasks)
}

func parseListFlags(args []string) (dto.GetTaskList, error) {
	newFlags := flag.NewFlagSet("list", flag.ContinueOnError)

	var sortedColumn string
	var status string
	var desc bool

	newFlags.StringVar(&sortedColumn, "s", "", "name of a column that will be sorted")
	newFlags.StringVar(&status, "f", "", "filter tasks by status (created, active, done, canceled)")
	newFlags.BoolVar(&desc, "desc", false, "sorting order")
	err := newFlags.Parse(args)

	if err != nil {
		return dto.GetTaskList{}, err
	}

	input := dto.GetTaskList{
		Status: status,
		SortBy: sortedColumn,
		Desc:   desc,
	}

	return input, nil
}
