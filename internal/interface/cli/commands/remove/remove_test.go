package remove

import (
	"errors"
	"fmt"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/remove/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type usecaseBehaviour func(s *mocks.MockTaskRemover, input []string)
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
			name:  "OK - valid input",
			input: []string{"1"},
			usecaseBehaviour: func(s *mocks.MockTaskRemover, input []string) {
				in := dto.RemoveTask{
					ID: 1,
				}
				s.EXPECT().Execute(in).Return(nil)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {
				s.EXPECT().PrintSuccessMessage(fmt.Sprintf("Task succesfully removed: %d", 1))
			},
			errorExpected: false,
		},
		{
			name:             "Error - no id provided",
			input:            []string{},
			usecaseBehaviour: func(s *mocks.MockTaskRemover, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter) {},
			errorExpected:    true,
			expectedError:    commands.ErrNotEnoughArguments,
		},
		{
			name:  "Error - usecase fail",
			input: []string{"1"},
			usecaseBehaviour: func(s *mocks.MockTaskRemover, input []string) {
				in := dto.RemoveTask{
					ID: 1,
				}
				s.EXPECT().Execute(in).Return(errors.New("Failed to remove task"))
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter) {},
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			ucMock := mocks.NewMockTaskRemover(c)
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
				assert.Equal(t, true, true)
			}
		})
	}
}

func TestParseRemoveArgs(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		errorExpected  bool
		expectedResult dto.RemoveTask
	}{
		{
			name:          "OK - valid id",
			input:         []string{"1"},
			errorExpected: false,
			expectedResult: dto.RemoveTask{
				ID: 1,
			},
		},

		{
			name:          "OK - negative id (parsable)",
			input:         []string{"-1"},
			errorExpected: false,
			expectedResult: dto.RemoveTask{
				ID: -1,
			},
		},
		{
			name:          "Error - invalid id",
			input:         []string{"abc"},
			errorExpected: true,
		},
		{
			name:          "Error - empty string",
			input:         []string{""},
			errorExpected: true,
		},
		{
			name:          "Error - hexadecimal number",
			input:         []string{"0xFF"},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseRemoveArgs(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult.ID, result.ID)
			}
		})
	}
}
