package writer

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestIsColumnNameValid(t *testing.T) {
	testCases := []struct {
		name            string
		input           string
		expectedContain bool
	}{
		{
			name:            "OK - valid name: id",
			input:           "id",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: name",
			input:           "name",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: description",
			input:           "description",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: created_at",
			input:           "created",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: updated_at",
			input:           "updated",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: status",
			input:           "status",
			expectedContain: true,
		},
		{
			name:            "OK - valid name: upper register",
			input:           "Status",
			expectedContain: true,
		},
		{
			name:            "Error - invalid name: invalid",
			input:           "invalid",
			expectedContain: false,
		},
		{
			name:            "Error - invalid name: empty",
			input:           "",
			expectedContain: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := isColumnNameValid(tc.input)

			assert.Equal(t, result, tc.expectedContain)
		})
	}
}

func TestPrintSuccessMessage(t *testing.T) {

	testCases := []struct {
		name           string
		input          string
		colored        bool
		expectedResult string
	}{
		{
			name:           "OK - colored",
			input:          "test message",
			colored:        true,
			expectedResult: "\x1b[32mtest message\x1b[0m\n",
		},
		{
			name:           "OK - no color",
			input:          "test message",
			colored:        false,
			expectedResult: "test message\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			oldNoColor := color.NoColor
			color.NoColor = !tc.colored

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			cfg := Config{
				MaxColumnWidth: 50,
				ExtraColumns:   []string{},
			}

			tw := New(cfg)
			tw.PrintSuccessMessage(tc.input)

			w.Close()
			os.Stdout = oldStdout
			color.NoColor = oldNoColor

			var buf bytes.Buffer
			io.Copy(&buf, r)
			actualOutpur := buf.String()

			assert.Equal(t, tc.expectedResult, actualOutpur)
		})
	}
}

func TestRenderTable(t *testing.T) {

	var (
		createdAt = time.Now().UTC()
		updatedAt = createdAt

		task1 = domain.Task{
			ID:          1,
			Name:        "task1",
			Description: nil,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Status:      domain.TaskStatusCreated,
		}
		task2 = domain.Task{
			ID:          2,
			Name:        "task2",
			Description: nil,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Status:      domain.TaskStatusActive,
		}
		task3 = domain.Task{
			ID:          3,
			Name:        "task3",
			Description: nil,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Status:      domain.TaskStatusCompleted,
		}
		task4 = domain.Task{
			ID:          4,
			Name:        "task4",
			Description: nil,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Status:      domain.TaskStatusCancelled,
		}
	)

	testCases := []struct {
		name              string
		input             []domain.Task
		extraColumns      []string
		colored           bool
		expectedContain   []string
		unexpectedContain []string
	}{
		{
			name:              "OK - no tasks",
			input:             []domain.Task{},
			extraColumns:      []string{},
			colored:           false,
			expectedContain:   []string{"No tasks yet..."},
			unexpectedContain: []string{"ID", "NAME", "DESCRIPTION", "CREATED", "UPDATED", "STATUS"},
		},
		{
			name:         "OK - all statuses colored",
			input:        []domain.Task{task1, task2, task3, task4},
			extraColumns: []string{},
			colored:      true,
			expectedContain: []string{
				color.HiGreenString(domain.StrCompleted),
				color.HiYellowString(domain.StrActive),
				color.HiBlueString(domain.StrCreated),
				color.HiRedString(domain.StrCancelled),
				"ID", "NAME", "STATUS",
			},
			unexpectedContain: []string{
				"No tasks yet...",
				"DESCRIPTION", "CREATED", "UPDATED",
			},
		},
		{
			name:         "OK - all statuses no color",
			input:        []domain.Task{task1, task2, task3, task4},
			extraColumns: []string{},
			colored:      false,
			expectedContain: []string{
				"ID", "NAME", "STATUS",
				domain.StrCompleted,
				domain.StrActive,
				domain.StrCreated,
				domain.StrCancelled,
			},
			unexpectedContain: []string{
				"No tasks yet...",
				"DESCRIPTION", "CREATED", "UPDATED",
				"\x1b",
			},
		},
		{
			name:  "OK - all extra columns",
			input: []domain.Task{task1, task2, task3, task4},
			extraColumns: []string{
				ValidDescriptionName,
				ValidCreatedName,
				ValidUpdatedName,
			},
			colored: false,
			expectedContain: []string{
				"ID", "NAME", "STATUS", "DESCRIPTION", "CREATED", "UPDATED",
			},
			unexpectedContain: []string{
				"No tasks yet...",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			oldNoColor := color.NoColor
			color.NoColor = !tc.colored

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			cfg := Config{
				MaxColumnWidth: 50,
				ExtraColumns:   tc.extraColumns,
			}

			tw := New(cfg)
			tw.RenderTable(tc.input)

			w.Close()
			os.Stdout = oldStdout
			color.NoColor = oldNoColor

			var buf bytes.Buffer
			io.Copy(&buf, r)
			actualOutpur := buf.String()

			for _, expectedString := range tc.expectedContain {
				assert.Contains(t, actualOutpur, expectedString)
			}
			for _, unexpectedString := range tc.unexpectedContain {
				assert.NotContains(t, actualOutpur, unexpectedString)
			}
		})
	}
}
