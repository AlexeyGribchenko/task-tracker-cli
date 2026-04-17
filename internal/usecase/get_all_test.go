package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskGetter)

	var (
		createdAt = time.Now().UTC()
		updatedAt = createdAt
	)

	testCases := []struct {
		name           string
		mockBehaviour  mockBehaviour
		errorExpected  bool
		expectedResult []domain.Task
	}{
		{
			name: "OK - not empty list",
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return([]domain.Task{
					{
						Name:        "task1",
						Description: nil,
						Status:      domain.TaskStatusCreated,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					},
					{
						Name:        "task2",
						Description: nil,
						Status:      domain.TaskStatusInProgress,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					},
				}, nil)
			},
			errorExpected: false,
			expectedResult: []domain.Task{
				{
					Name:        "task1",
					Description: nil,
					Status:      domain.TaskStatusCreated,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				},
				{
					Name:        "task2",
					Description: nil,
					Status:      domain.TaskStatusInProgress,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				},
			},
		},
		{
			name: "OK - empty list",
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return([]domain.Task{}, nil)
			},
			errorExpected:  false,
			expectedResult: []domain.Task{},
		},
		{
			name: "Error - database error",
			mockBehaviour: func(s *mocks.MockTaskGetter) {
				s.EXPECT().GetTasks().Return(nil, errors.New("database connection error"))
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			mockRepo := mocks.NewMockTaskGetter(c)
			tc.mockBehaviour(mockRepo)

			getUC := NewGetTasksUseCase(mockRepo)

			results, err := getUC.Execute()

			if tc.errorExpected {
				assert.Error(t, err)
				assert.Nil(t, results)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, len(results), len(tc.expectedResult))
				assert.ElementsMatch(t, results, tc.expectedResult)
			}
		})
	}
}
