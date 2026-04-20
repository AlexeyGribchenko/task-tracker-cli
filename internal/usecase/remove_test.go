package usecase

import (
	"fmt"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskRemover, input dto.RemoveTask)

	testCases := []struct {
		name          string
		input         dto.RemoveTask
		mockBehaviour mockBehaviour
		errorExpected bool
		errorIs       bool
		expectedError error
	}{
		{
			name:  "OK - valid delete",
			input: dto.RemoveTask{ID: 1},
			mockBehaviour: func(s *mocks.MockTaskRemover, input dto.RemoveTask) {
				s.EXPECT().RemoveTask(input.ID).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "Error - error from storage",
			input: dto.RemoveTask{ID: 1},
			mockBehaviour: func(s *mocks.MockTaskRemover, input dto.RemoveTask) {
				s.EXPECT().RemoveTask(input.ID).Return(fmt.Errorf("storage: task with id 2 not found"))
			},
			errorExpected: true,
		},
		{
			name:          "Error - invalid task ID",
			input:         dto.RemoveTask{ID: -1},
			mockBehaviour: func(s *mocks.MockTaskRemover, input dto.RemoveTask) {},
			errorExpected: true,
			errorIs:       true,
			expectedError: ErrInvalidTaskID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			repo := mocks.NewMockTaskRemover(c)

			uc := NewRemoveTaskUseCase(repo)

			tc.mockBehaviour(repo, tc.input)

			err := uc.Execute(tc.input)

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
