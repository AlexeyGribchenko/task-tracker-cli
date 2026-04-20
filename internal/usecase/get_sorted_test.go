package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetSorted(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted)

	tasks := make([]domain.Task, 3)
	tasks[0] = domain.Task{
		ID:          1,
		Name:        "task3",
		Description: nil,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Status:      domain.TaskStatusCreated,
	}
	tasks[1] = domain.Task{
		ID:          2,
		Name:        "task2",
		Description: nil,
		CreatedAt:   time.Now().Add(time.Second).UTC(),
		UpdatedAt:   time.Now().Add(5 * time.Second).UTC(),
		Status:      domain.TaskStatusCompleted,
	}
	tasks[2] = domain.Task{
		ID:          3,
		Name:        "task1",
		Description: nil,
		CreatedAt:   time.Now().Add(2 * time.Second).UTC(),
		UpdatedAt:   time.Now().Add(2 * time.Second).UTC(),
		Status:      domain.TaskStatusInProgress,
	}

	testCases := []struct {
		name           string
		input          dto.GetTasksSorted
		mockBehaviour  mockBehaviour
		errorExpected  bool
		errorIs        bool
		expectedError  error
		expectedResult []domain.Task
	}{
		{
			name:  "OK - sort status",
			input: dto.GetTasksSorted{ColumnSorted: "status"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {
				s.EXPECT().GetSorted(domain.ColumnName(input.ColumnSorted)).Return(
					[]domain.Task{tasks[1], tasks[0], tasks[2]},
					nil,
				)
			},
			expectedResult: []domain.Task{tasks[1], tasks[0], tasks[2]},
			errorExpected:  false,
		},
		{
			name:  "OK - sort name",
			input: dto.GetTasksSorted{ColumnSorted: "name"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {
				s.EXPECT().GetSorted(domain.ColumnName(input.ColumnSorted)).Return(
					[]domain.Task{tasks[2], tasks[1], tasks[0]},
					nil,
				)
			},
			expectedResult: []domain.Task{tasks[2], tasks[1], tasks[0]},
			errorExpected:  false,
		},
		{
			name:  "OK - sort created_at",
			input: dto.GetTasksSorted{ColumnSorted: "created_at"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {
				s.EXPECT().GetSorted(domain.ColumnName(input.ColumnSorted)).Return(
					[]domain.Task{tasks[0], tasks[1], tasks[2]},
					nil,
				)
			},
			expectedResult: []domain.Task{tasks[0], tasks[1], tasks[2]},
			errorExpected:  false,
		},
		{
			name:  "OK - sort updated_at",
			input: dto.GetTasksSorted{ColumnSorted: "updated_at"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {
				s.EXPECT().GetSorted(domain.ColumnName(input.ColumnSorted)).Return(
					[]domain.Task{tasks[0], tasks[2], tasks[1]},
					nil,
				)
			},
			expectedResult: []domain.Task{tasks[0], tasks[2], tasks[1]},
			errorExpected:  false,
		},
		{
			name:          "Error - invalid column name",
			input:         dto.GetTasksSorted{ColumnSorted: "invalid_name"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {},
			errorExpected: true,
			errorIs:       true,
			expectedError: domain.ErrInvalidColumnName,
		},
		{
			name:  "Error - database error",
			input: dto.GetTasksSorted{ColumnSorted: "status"},
			mockBehaviour: func(s *mocks.MockTaskSortedGetter, input dto.GetTasksSorted) {
				s.EXPECT().GetSorted(domain.ColumnName(input.ColumnSorted)).Return(
					nil,
					errors.New("database error"),
				)
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			repo := mocks.NewMockTaskSortedGetter(c)

			uc := NewGetTasksSorted(repo)

			tc.mockBehaviour(repo, tc.input)

			tasks, err := uc.Execute(tc.input)
			if tc.errorExpected {
				assert.Error(t, err)
				if tc.errorIs {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NotNil(t, tasks)
				assert.Equal(t, len(tasks), len(tc.expectedResult))

				for i := range tasks {
					assert.Equal(t, tasks[i].ID, tc.expectedResult[i].ID)
				}
			}
		})
	}
}
