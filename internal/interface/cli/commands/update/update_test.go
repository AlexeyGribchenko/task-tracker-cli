package update

import (
	"errors"
	"fmt"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/update/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	type usecaseBehaviour func(s *mocks.MockTaskUpdater, input []string)
	type writerBehaviour func(s *mocks.MockSuccesWriter)

	testCases := []struct {
		name             string
		input            []string
		usecaseBehaviour usecaseBehaviour
		writerBehaviour  writerBehaviour
		errorExpected    bool
		expectedError    error
	}{
		{
			name:  "OK - valid id and status",
			input: []string{"1", "completed"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {
				in := dto.UpdateTask{
					ID:     1,
					Status: input[1],
				}
				s.EXPECT().Execute(in).Return(nil)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {
				s.EXPECT().PrintSuccessMessage("Task status successuly updated: 1")
			},
			errorExpected: false,
		},
		{
			name:             "Error - invalid task ID",
			input:            []string{"invalid", "done"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter) {},
			errorExpected:    true,
		},
		{
			name:  "Error - task not found",
			input: []string{"999", "created"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {
				in := dto.UpdateTask{
					ID:     999,
					Status: input[1],
				}
				s.EXPECT().Execute(in).Return(errors.New("Task not found"))
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {},
			errorExpected:   true,
		},
		{
			name:  "Error - invalid status value",
			input: []string{"1", "invalid"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {
				in := dto.UpdateTask{
					ID:     1,
					Status: input[1],
				}
				s.EXPECT().Execute(in).Return(
					fmt.Errorf("Failed to update task status"),
				)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {},
			errorExpected:   true,
		},
		{
			name:             "Error - not enough arguments (empty slice)",
			input:            []string{},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter) {},
			errorExpected:    true,
		},
		{
			name:             "Error - not enough arguments (only id)",
			input:            []string{"1"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter) {},
			errorExpected:    true,
		},
		{
			name:  "Error - usecase returns unexpected error",
			input: []string{"1", "created"},
			usecaseBehaviour: func(s *mocks.MockTaskUpdater, input []string) {
				in := dto.UpdateTask{
					ID:     1,
					Status: input[1],
				}
				s.EXPECT().Execute(in).Return(fmt.Errorf("database connection error"))
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {},
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			ucMock := mocks.NewMockTaskUpdater(c)
			wrMock := mocks.NewMockSuccesWriter(c)

			tc.usecaseBehaviour(ucMock, tc.input)
			tc.writerBehaviour(wrMock)

			cmd := New(ucMock, wrMock)

			err := cmd.Execute(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				if tc.expectedError != nil {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseUpdateArgs(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		errorExpected  bool
		expectedError  error
		expectedResult dto.UpdateTask
	}{
		{
			name:          "OK - valid id and status",
			input:         []string{"1", "completed"},
			errorExpected: false,
			expectedResult: dto.UpdateTask{
				ID:     1,
				Status: "completed",
			},
		},
		{
			name:          "OK - zero ID",
			input:         []string{"0", "in-progress"},
			errorExpected: false,
			expectedResult: dto.UpdateTask{
				ID:     0,
				Status: "in-progress",
			},
		},
		{
			name:          "Error - invalid ID (string)",
			input:         []string{"abc", "done"},
			errorExpected: true,
		},
		{
			name:          "Error - negative ID",
			input:         []string{"-1", "done"},
			errorExpected: false,
			expectedResult: dto.UpdateTask{
				ID:     -1,
				Status: "done",
			},
		},
		{
			name:          "Error - empty string as ID",
			input:         []string{"", "done"},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseUpdateArgs(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				if tc.expectedError != nil {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult.ID, result.ID)
				assert.Equal(t, tc.expectedResult.Status, result.Status)
			}
		})
	}
}
