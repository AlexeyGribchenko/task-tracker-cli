package usecase

import (
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskGetter)

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
		name           string
		input          dto.GetTaskList
		mockBehaviour  mockBehaviour
		errorExpected  bool
		errorIs        bool
		expectedError  error
		expectedResult []domain.Task
	}{
		{
			name:  "OK - not empty list",
			input: dto.GetTaskList{},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return([]domain.Task{
					task1, task2,
				}, nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1, task2,
			},
		},
		{
			name:  "OK - sort id",
			input: dto.GetTaskList{SortBy: domain.ColumnId},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1, task2, task3, task4,
			},
		},
		{
			name:  "OK - sort name",
			input: dto.GetTaskList{SortBy: domain.ColumnName},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1, task2, task3, task4,
			},
		},
		{
			name:  "OK - sort description",
			input: dto.GetTaskList{SortBy: domain.ColumnDescription},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task4, task2, task1, task3,
			},
		},
		{
			name:  "OK - sort created",
			input: dto.GetTaskList{SortBy: domain.ColumnCreatedAt},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1, task3, task2, task4,
			},
		},
		{
			name:  "OK - sort updated",
			input: dto.GetTaskList{SortBy: domain.ColumnUpdatedAt},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1, task3, task4, task2,
			},
		},
		{
			name:  "OK - sort status",
			input: dto.GetTaskList{SortBy: domain.ColumnStatus},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task3, task4, task2, task1,
			},
		},
		{
			name:  "OK - sort name decs",
			input: dto.GetTaskList{SortBy: domain.ColumnName, Desc: true},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task4, task3, task2, task1,
			},
		},
		{
			name:  "OK - filter active",
			input: dto.GetTaskList{Status: domain.StrActive},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected:  false,
			expectedResult: []domain.Task{task3},
		},
		{
			name:  "OK - filter cancelled",
			input: dto.GetTaskList{Status: domain.StrCancelled},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected:  false,
			expectedResult: []domain.Task{task4},
		},
		{
			name:  "OK - filter completed",
			input: dto.GetTaskList{Status: domain.StrCompleted},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task2,
			},
		},
		{
			name:  "OK - filter created",
			input: dto.GetTaskList{Status: domain.StrCreated},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(slices.Clone(taskList), nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				task1,
			},
		},
		{
			name:  "OK - empty list",
			input: dto.GetTaskList{},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return([]domain.Task{}, nil)
			},
			errorExpected:  false,
			expectedResult: []domain.Task{},
		},
		{
			name:  "Error - database error",
			input: dto.GetTaskList{},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(nil, errors.New("database connection error"))
			},
			errorExpected: true,
		},
		{
			name:  "Error - sort: invalid column name",
			input: dto.GetTaskList{SortBy: "invalid column name"},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(taskList, nil)
			},
			errorExpected: true,
			errorIs:       true,
			expectedError: domain.ErrInvalidColumnName,
		},
		{
			name:  "Error - filter: invalid status name",
			input: dto.GetTaskList{Status: "invalid status"},
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(taskList, nil)
			},
			errorExpected: true,
			errorIs:       true,
			expectedError: domain.ErrInvalidTaskStatus,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			mockRepo := mocks.NewMockTaskGetter(c)
			tc.mockBehaviour(mockRepo)

			getUC := NewGetTasksUseCase(mockRepo)

			results, err := getUC.Execute(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				assert.Nil(t, results)
				if tc.errorIs {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, len(results), len(tc.expectedResult))
				assert.ElementsMatch(t, results, tc.expectedResult)
			}
		})
	}
}
