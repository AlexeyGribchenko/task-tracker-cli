package usecase

import (
	"errors"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskRepository, input dto.UpdateTask)

	testCases := []struct {
		name          string
		input         dto.UpdateTask
		mockBehaviour mockBehaviour
		errorExpected bool
		errorIs       bool
		expectedError error
	}{
		{
			name:  "OK - valid input: in_progress status",
			input: dto.UpdateTask{ID: 1, Status: "in_progress"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {
				s.EXPECT().UpdateTaskStatus(input.ID, domain.TaskStatusInProgress).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "OK - valid input: completed status",
			input: dto.UpdateTask{ID: 1, Status: "completed"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {
				s.EXPECT().UpdateTaskStatus(input.ID, domain.TaskStatusCompleted).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "OK - valid input: canceled status",
			input: dto.UpdateTask{ID: 1, Status: "cancelled"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {
				s.EXPECT().UpdateTaskStatus(input.ID, domain.TaskStatusCancelled).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:          "Error - invalid input: wrong ID",
			input:         dto.UpdateTask{ID: -1, Status: "in_progress"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {},
			errorExpected: true,
			errorIs:       true,
			expectedError: ErrInvalidTaskID,
		},
		{
			name:          "Error - invalid input: wrong task status",
			input:         dto.UpdateTask{ID: 1, Status: "wrong status"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {},
			errorExpected: true,
			errorIs:       true,
			expectedError: ErrInvalidTaskStatus,
		},
		{
			name:  "Error - database error",
			input: dto.UpdateTask{ID: 1, Status: "in_progress"},
			mockBehaviour: func(s *mocks.MockTaskRepository, input dto.UpdateTask) {
				s.EXPECT().UpdateTaskStatus(input.ID, domain.TaskStatusInProgress).
					Return(errors.New("database"))
			},
			errorExpected: true,
			errorIs:       false,
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			repo := mocks.NewMockTaskRepository(c)

			updateUC := NewUpdateTaskUseCase(repo)

			tc.mockBehaviour(repo, tc.input)

			err := updateUC.Execute(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				if tc.errorIs {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
