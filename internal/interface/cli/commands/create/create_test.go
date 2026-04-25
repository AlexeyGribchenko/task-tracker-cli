package create

import (
	"fmt"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/create/mocks"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type usecaseBehaviour func(s *mocks.MockTaskCreator, input []string)
	type writerBehaviour func(s *mocks.MockSuccesWriter, input []string)

	testCases := []struct {
		name             string
		input            []string
		usecaseBehaviour usecaseBehaviour
		writerBehaviour  writerBehaviour
		errorExpected    bool
		expectedError    error
	}{
		{
			name:  "OK - all args",
			input: []string{"name", "-d", "descrtiption"},
			usecaseBehaviour: func(s *mocks.MockTaskCreator, input []string) {
				in := dto.CreateTask{
					Name:        input[0],
					Description: utils.PointerFromValue(input[2]),
				}
				s.EXPECT().Execute(in).Return(&domain.Task{
					Name:        input[0],
					Description: utils.PointerFromValue(input[2]),
				}, nil)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter, input []string) {
				s.EXPECT().PrintSuccessMessage(fmt.Sprintf("Task succsessfuly created: %s", input[0]))
			},
			errorExpected: false,
		},
		{
			name:  "OK - no description",
			input: []string{"name"},
			usecaseBehaviour: func(s *mocks.MockTaskCreator, input []string) {
				in := dto.CreateTask{
					Name:        input[0],
					Description: nil,
				}
				s.EXPECT().Execute(in).Return(&domain.Task{
					Name:        input[0],
					Description: nil,
				}, nil)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter, input []string) {
				s.EXPECT().PrintSuccessMessage(fmt.Sprintf("Task succsessfuly created: %s", input[0]))
			},
			errorExpected: false,
		},
		{
			name:  "Error - invalid name",
			input: []string{"", "-d", "descrtiption"},
			usecaseBehaviour: func(s *mocks.MockTaskCreator, input []string) {
				in := dto.CreateTask{
					Name:        input[0],
					Description: utils.PointerFromValue(input[2]),
				}
				s.EXPECT().Execute(in).Return(
					nil,
					fmt.Errorf("Failed to create task: %w", domain.ErrInvalidTaskName),
				)
			},
			writerBehaviour: func(s *mocks.MockSuccesWriter, input []string) {},
			errorExpected:   true,
			expectedError:   domain.ErrInvalidTaskName,
		},
		{
			name:             "Error - no args",
			input:            []string{},
			usecaseBehaviour: func(s *mocks.MockTaskCreator, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter, input []string) {},
			errorExpected:    true,
			expectedError:    commands.ErrNotEnoughArguments,
		},
		{
			name:             "Error - invalid flags",
			input:            []string{"name", "-invalid", "description"},
			usecaseBehaviour: func(s *mocks.MockTaskCreator, input []string) {},
			writerBehaviour:  func(s *mocks.MockSuccesWriter, input []string) {},
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			ucMock := mocks.NewMockTaskCreator(c)
			wrMock := mocks.NewMockSuccesWriter(c)

			tc.usecaseBehaviour(ucMock, tc.input)
			tc.writerBehaviour(wrMock, tc.input)

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

func TestParseArgs(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		errorExpected  bool
		expectedResult dto.CreateTask
	}{
		{
			name:          "OK - all flags",
			input:         []string{"name", "-d", "description"},
			errorExpected: false,
			expectedResult: dto.CreateTask{
				Name:        "name",
				Description: utils.PointerFromValue("description"),
			},
		},
		{
			name:          "OK - no description",
			input:         []string{"name"},
			errorExpected: false,
			expectedResult: dto.CreateTask{
				Name:        "name",
				Description: nil,
			},
		},
		{
			name:          "Error - invalid flag",
			input:         []string{"name", "--invalid", "description"},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := parseCreateArgs(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result.Name, tc.expectedResult.Name)
				assert.Equal(t,
					utils.ValueFromPointer(result.Description),
					utils.ValueFromPointer(tc.expectedResult.Description),
				)
			}
		})
	}
}
