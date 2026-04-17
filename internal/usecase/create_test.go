package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const timeDelta = 10 * time.Millisecond

type taskMatcher struct {
	expected domain.Task
}

func (t taskMatcher) Matches(x any) bool {
	actual, ok := x.(domain.Task)
	if !ok {
		return false
	}

	return t.expected.Name == actual.Name &&
		reflect.DeepEqual(t.expected.Description, actual.Description) &&
		t.expected.Status == actual.Status &&
		t.expected.CreatedAt.Sub(actual.CreatedAt) < timeDelta &&
		t.expected.UpdatedAt.Sub(actual.UpdatedAt) < timeDelta
}

func (t taskMatcher) String() string {
	return "popa"
}

func TestCreate(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskCreator, input domain.Task)

	var (
		ptrDescription = utils.PointerFromValue("description")
		createdAt      = time.Now().UTC()
		updatedAt      = createdAt
	)

	testCases := []struct {
		name           string
		input          dto.CreateTask
		inputDB        domain.Task
		mockBehaviour  mockBehaviour
		errorExpected  bool
		errorIs        bool
		expectedError  error
		expectedResult *domain.Task
	}{
		{
			name:  "OK - create task with all fields",
			input: dto.CreateTask{Name: "name", Description: ptrDescription},
			inputDB: domain.Task{
				Name:        "name",
				Description: ptrDescription,
				Status:      domain.TaskStatusCreated,
				UpdatedAt:   updatedAt,
				CreatedAt:   createdAt,
			},
			mockBehaviour: func(s *mocks.MockTaskCreator, input domain.Task) {
				s.EXPECT().CreateTask(taskMatcher{expected: input}).Return(
					&domain.Task{
						ID:          1,
						Name:        "name",
						Description: ptrDescription,
						Status:      domain.TaskStatusCreated,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					}, nil,
				)
			},
			errorExpected: false,
			expectedResult: &domain.Task{
				ID:          1,
				Name:        "name",
				Description: ptrDescription,
				Status:      domain.TaskStatusCreated,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
		},
		{
			name:  "OK - nil description",
			input: dto.CreateTask{Name: "name", Description: nil},
			inputDB: domain.Task{
				Name:        "name",
				Description: nil,
				Status:      domain.TaskStatusCreated,
				UpdatedAt:   updatedAt,
				CreatedAt:   createdAt,
			},
			mockBehaviour: func(s *mocks.MockTaskCreator, input domain.Task) {
				s.EXPECT().CreateTask(taskMatcher{expected: input}).Return(
					&domain.Task{
						ID:          1,
						Name:        "name",
						Description: nil,
						Status:      domain.TaskStatusCreated,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					}, nil,
				)
			},
			errorExpected: false,
			expectedResult: &domain.Task{
				ID:          1,
				Name:        "name",
				Description: nil,
				Status:      domain.TaskStatusCreated,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
		},
		{
			name:          "Error - empty name",
			input:         dto.CreateTask{Name: "", Description: nil},
			mockBehaviour: func(s *mocks.MockTaskCreator, input domain.Task) {},
			errorExpected: true,
			errorIs:       true,
			expectedError: domain.ErrInvalidTaskName,
		},
		{
			name:  "Error - Database error",
			input: dto.CreateTask{Name: "name", Description: nil},
			inputDB: domain.Task{
				Name:        "name",
				Description: nil,
				Status:      domain.TaskStatusCreated,
				UpdatedAt:   updatedAt,
				CreatedAt:   createdAt,
			},
			mockBehaviour: func(s *mocks.MockTaskCreator, input domain.Task) {
				s.EXPECT().CreateTask(taskMatcher{expected: input}).Return(nil, errors.New("database connection error"))
			},
			errorExpected: true,
			errorIs:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			mockRepo := mocks.NewMockTaskCreator(c)

			tc.mockBehaviour(mockRepo, tc.inputDB)

			uc := NewCreateTaskUseCase(mockRepo)

			task, err := uc.Execute(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				assert.Nil(t, task)
				if tc.errorIs {
					assert.ErrorIs(t, err, tc.expectedError)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
				assert.Equal(t, task.ID, tc.expectedResult.ID)
				assert.Equal(t, task.Name, tc.expectedResult.Name)
				assert.Equal(t, task.Description, tc.expectedResult.Description)
				assert.Equal(t, task.Status, tc.expectedResult.Status)
				assert.NotZero(t, task.CreatedAt)
				assert.NotZero(t, task.UpdatedAt)
				assert.WithinDuration(t, task.CreatedAt, tc.expectedResult.CreatedAt, timeDelta)
				assert.WithinDuration(t, task.UpdatedAt, tc.expectedResult.UpdatedAt, timeDelta)
			}
		})
	}
}
