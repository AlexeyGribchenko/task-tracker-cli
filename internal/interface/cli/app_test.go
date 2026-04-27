package cli

import (
	"os"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/mocks"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {

	type mockBehaviour func(s *mocks.MockRepository)

	var (
		createdAt = time.Now().UTC()
		updatedAt = createdAt

		task = &domain.Task{
			ID:          1,
			Name:        "test",
			Description: utils.PointerFromValue("description"),
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Status:      domain.TaskStatusCreated,
		}
	)

	testCases := []struct {
		name          string
		input         []string
		mockBehaviour mockBehaviour
		errorExpected bool
	}{
		{
			name: "OK - list (all flags)",
			input: []string{
				"appname", CommandGetTaskList, "-s", "name", "-f", "active", "-desc",
			},
			mockBehaviour: func(s *mocks.MockRepository) {
				s.EXPECT().GetTasks().Return([]domain.Task{*task}, nil)
			},
			errorExpected: false,
		},
		{
			name: "OK - add (all flags)",
			input: []string{
				"appname", CommandCreateTask, "test", "-d", "description",
			},
			mockBehaviour: func(s *mocks.MockRepository) {
				s.EXPECT().CreateTask(gomock.Any()).Return(task, nil)
			},
			errorExpected: false,
		},
		{
			name: "OK - status",
			input: []string{
				"appname", CommandSetTaskStatus, "1", "done",
			},
			mockBehaviour: func(s *mocks.MockRepository) {
				s.EXPECT().UpdateTaskStatus(1, domain.TaskStatusCompleted).Return(nil)
			},
			errorExpected: false,
		},
		{
			name: "OK - remove",
			input: []string{
				"appname", CommandRemoveTask, "1",
			},
			mockBehaviour: func(s *mocks.MockRepository) {
				s.EXPECT().RemoveTask(1).Return(nil)
			},
			errorExpected: false,
		},
		{
			name: "Error - not enough arguments",
			input: []string{
				"appname",
			},
			mockBehaviour: func(s *mocks.MockRepository) {},
			errorExpected: true,
		},
		{
			name: "Error - invalid flag",
			input: []string{
				"appname", "list", "-invalid",
			},
			mockBehaviour: func(s *mocks.MockRepository) {},
			errorExpected: true,
		},
		{
			name: "Error - invalid command",
			input: []string{
				"appname", "invalid",
			},
			mockBehaviour: func(s *mocks.MockRepository) {},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			repo := mocks.NewMockRepository(c)
			wr := writer.New(writer.Config{
				MaxColumnWidth: 50,
				ExtraColumns:   []string{},
			})

			tc.mockBehaviour(repo)

			oldArgs := os.Args
			os.Args = tc.input

			app := New(repo, *wr)

			err := app.Run()

			os.Args = oldArgs

			if tc.errorExpected {
				assert.Error(t, err)
			}
		})
	}
}
