package domain

import (
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseStatus(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		errorExpected  bool
		expectedError  error
		expectedStatus TaskStatus
	}{
		{
			name:           "OK - valid task status (in progress)",
			input:          "in_progress",
			errorExpected:  false,
			expectedStatus: TaskStatusInProgress,
		},
		{
			name:           "OK - valid task status (created)",
			input:          "created",
			errorExpected:  false,
			expectedStatus: TaskStatusCreated,
		},
		{
			name:           "OK - valid task status (completed)",
			input:          "completed",
			errorExpected:  false,
			expectedStatus: TaskStatusCompleted,
		},
		{
			name:           "OK - valid task status (cancelled)",
			input:          "cancelled",
			errorExpected:  false,
			expectedStatus: TaskStatusCancelled,
		},
		{
			name:          "Error - invalid task status",
			input:         "invalid",
			errorExpected: true,
			expectedError: ErrInvalidTaskStatus,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			status, err := ParseStatus(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, status, tc.expectedStatus)
			}
		})
	}
}

func TestNewTask(t *testing.T) {

	var (
		inputDescription    = utils.PointerFromValue("test")
		nilInputDescription = (*string)(nil)
		createdAt           = time.Now().UTC()
		updatedAt           = time.Now().UTC()
		deltaTime           = 10 * time.Millisecond
	)

	testCases := []struct {
		name             string
		inputName        string
		inputDescription *string
		expectedTask     *Task
		errorExpected    bool
		expectedError    error
	}{
		{
			name:             "OK - all fields",
			inputName:        "test",
			inputDescription: inputDescription,
			expectedTask: &Task{
				Name:        "test",
				Description: inputDescription,
				Status:      TaskStatusCreated,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
			errorExpected: false,
		},
		{
			name:             "OK - nil description",
			inputName:        "test",
			inputDescription: nilInputDescription,
			expectedTask: &Task{
				Name:        "test",
				Description: nil,
				Status:      TaskStatusCreated,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
			errorExpected: false,
		},
		{
			name:             "Error - empty task name",
			inputName:        "",
			inputDescription: inputDescription,
			errorExpected:    true,
			expectedError:    ErrInvalidTaskName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			task, err := NewTask(tc.inputName, tc.inputDescription)

			if tc.errorExpected {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, task.Name, tc.expectedTask.Name)
				if tc.expectedTask.Description == nil {
					assert.Nil(t, task.Description)
				} else {
					assert.Equal(t, *task.Description, *tc.expectedTask.Description)
				}
				assert.Equal(t, task.Status, tc.expectedTask.Status)
				assert.WithinDuration(t, task.CreatedAt, tc.expectedTask.CreatedAt, deltaTime)
				assert.WithinDuration(t, task.UpdatedAt, tc.expectedTask.UpdatedAt, deltaTime)
			}
		})
	}
}

func TestTaskStatusString(t *testing.T) {
	testCases := []struct {
		name           string
		status         TaskStatus
		expectedOutput string
	}{
		{
			name:           "OK - Status 'in_progress'",
			status:         TaskStatusInProgress,
			expectedOutput: "in_progress",
		},
		{
			name:           "OK - Status 'created'",
			status:         TaskStatusCreated,
			expectedOutput: "created",
		},
		{
			name:           "OK - Status 'completed'",
			status:         TaskStatusCompleted,
			expectedOutput: "completed",
		},
		{
			name:           "OK - Status 'cancelled'",
			status:         TaskStatusCancelled,
			expectedOutput: "cancelled",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			output := tc.status.String()

			assert.Equal(t, output, tc.expectedOutput)
		})
	}
}
