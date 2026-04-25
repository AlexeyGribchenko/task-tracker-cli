package list

import (
	"errors"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli/commands/list/mocks"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	type usecaseBehaviour func(s *mocks.MockTasksGetter, input []string)
	type writerBehaviuor func(s *mocks.MockTaskWriter)

	var (
		createdAt = time.Now().UTC()
		updatedAt = createdAt

		task1 = domain.Task{
			ID:          1,
			Name:        "task1",
			Description: utils.PointerFromValue("ccc"),
			Status:      domain.TaskStatusCreated,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		task2 = domain.Task{
			ID:          2,
			Name:        "task2",
			Description: utils.PointerFromValue("bbb"),
			Status:      domain.TaskStatusCompleted,
			CreatedAt:   createdAt.Add(2 * time.Hour),
			UpdatedAt:   updatedAt.Add(8 * time.Hour),
		}
		task3 = domain.Task{
			ID:          3,
			Name:        "task3",
			Description: utils.PointerFromValue("ddd"),
			Status:      domain.TaskStatusActive,
			CreatedAt:   createdAt.Add(1 * time.Hour),
			UpdatedAt:   updatedAt.Add(1 * time.Hour),
		}
		task4 = domain.Task{
			ID:          4,
			Name:        "task4",
			Description: utils.PointerFromValue("aaa"),
			Status:      domain.TaskStatusCancelled,
			CreatedAt:   createdAt.Add(4 * time.Hour),
			UpdatedAt:   updatedAt.Add(5 * time.Hour),
		}
		taskList = []domain.Task{task1, task3, task2, task4}
	)

	testCases := []struct {
		name             string
		input            []string
		usecaseBehaviour usecaseBehaviour
		writerBehaviour  writerBehaviuor
		errorExpected    bool
		expectedError    error
	}{
		{
			name:  "OK - simple list (without flags)",
			input: []string{},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{}
				s.EXPECT().Execute(in).Return(taskList, nil)
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {
				s.EXPECT().RenderTable(taskList).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "OK - sort desc",
			input: []string{"-desc"},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{
					Desc: true,
				}
				s.EXPECT().Execute(in).Return([]domain.Task{
					task4, task3, task2, task1,
				}, nil)
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {
				s.EXPECT().RenderTable([]domain.Task{
					task4, task3, task2, task1,
				}).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "OK - filter status",
			input: []string{"-f", "active"},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{
					Status: input[1],
				}
				s.EXPECT().Execute(in).Return([]domain.Task{
					task3,
				}, nil)
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {
				s.EXPECT().RenderTable([]domain.Task{
					task3,
				}).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:  "OK - filter all args",
			input: []string{"-s", "status", "-f", "active", "-desc"},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{
					Status: input[3],
					Desc:   true,
					SortBy: input[1],
				}
				s.EXPECT().Execute(in).Return([]domain.Task{
					task3,
				}, nil)
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {
				s.EXPECT().RenderTable([]domain.Task{
					task3,
				}).Return(nil)
			},
			errorExpected: false,
		},
		{
			name:             "Error - invalid flag",
			input:            []string{"-invalid", "value"},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {},
			writerBehaviour:  func(s *mocks.MockTaskWriter) {},
			errorExpected:    true,
		},
		{
			name:  "Error - usecase fail",
			input: []string{},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{}
				s.EXPECT().Execute(in).Return(nil, errors.New("Failed to get tasks"))
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {},
			errorExpected:   true,
		},
		{
			name:  "Error - writer fail",
			input: []string{},
			usecaseBehaviour: func(s *mocks.MockTasksGetter, input []string) {
				in := dto.GetTaskList{}
				s.EXPECT().Execute(in).Return(taskList, nil)
			},
			writerBehaviour: func(s *mocks.MockTaskWriter) {
				s.EXPECT().RenderTable(taskList).Return(errors.New("Failed to render table"))
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			ucMock := mocks.NewMockTasksGetter(c)
			wrMock := mocks.NewMockTaskWriter(c)

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

func TestParseListFlags(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		errorExpected  bool
		expectedResult dto.GetTaskList
	}{
		{
			name:          "OK - all flags",
			input:         []string{"-s", "name", "-f", "active", "-desc"},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "active",
				SortBy: "name",
				Desc:   true,
			},
		},
		{
			name:          "OK - no flags",
			input:         []string{},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "",
				SortBy: "",
				Desc:   false,
			},
		},
		{
			name:          "OK - only status filter",
			input:         []string{"-f", "completed"},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "completed",
				SortBy: "",
				Desc:   false,
			},
		},
		{
			name:          "OK - only sort column",
			input:         []string{"-s", "created"},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "",
				SortBy: "created",
				Desc:   false,
			},
		},
		{
			name:          "OK - sort with desc order",
			input:         []string{"-s", "name", "-desc"},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "",
				SortBy: "name",
				Desc:   true,
			},
		},
		{
			name:          "OK - only desc",
			input:         []string{"-desc"},
			errorExpected: false,
			expectedResult: dto.GetTaskList{
				Status: "",
				SortBy: "",
				Desc:   true,
			},
		},
		{
			name:          "Error - invalid flag",
			input:         []string{"-invalid", "value"},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseListFlags(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult.Status, result.Status)
				assert.Equal(t, tc.expectedResult.SortBy, result.SortBy)
				assert.Equal(t, tc.expectedResult.Desc, result.Desc)
			}
		})
	}
}
