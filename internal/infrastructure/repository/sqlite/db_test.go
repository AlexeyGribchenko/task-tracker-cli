package sqlite

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/domain"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	db      *sql.DB
	storage *Storage
}

func (s *TaskRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(s.T(), err)

	s.db = db
	s.storage = &Storage{db: db}

	s.createTables()
}

func (s *TaskRepositoryTestSuite) TearDownTest() {
	s.db.Close()
}

func (s *TaskRepositoryTestSuite) createTables() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			status TEXT,
			created_at DATE,
			updated_at DATE
		);

		CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);

		INSERT INTO tasks(name, description, status, created_at, updated_at)
		VALUES
		('test1', '', 'done', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
		('test2', 'description', 'active', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)
	assert.NoError(s.T(), err)
}

func TestTaskStorageTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (s *TaskRepositoryTestSuite) TestCreateTask() {
	testCases := []struct {
		name string
		task domain.Task
	}{
		{
			name: "OK - task without description",
			task: domain.Task{
				Name:        "test",
				Description: nil,
				Status:      domain.TaskStatusCreated,
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
			},
		},
		{
			name: "OK - task with description",
			task: domain.Task{
				Name:        "test",
				Description: utils.PointerFromValue("description"),
				Status:      domain.TaskStatusCreated,
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
			},
		},
	}

	nextTaskID := 3

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			createdTask, err := s.storage.CreateTask(tc.task)

			assert.NoError(s.T(), err)

			assert.Equal(s.T(), nextTaskID, createdTask.ID)
			assert.Equal(s.T(), tc.task.Name, createdTask.Name)
			assert.Equal(s.T(),
				utils.ValueFromPointer(tc.task.Description),
				utils.ValueFromPointer(createdTask.Description),
			)
			assert.Equal(s.T(), tc.task.Status, createdTask.Status)
			assert.WithinDuration(s.T(),
				tc.task.CreatedAt, createdTask.CreatedAt, time.Millisecond,
			)
			assert.WithinDuration(s.T(),
				tc.task.UpdatedAt, createdTask.UpdatedAt, time.Millisecond,
			)
		})
		nextTaskID++
	}
}

func (s *TaskRepositoryTestSuite) TestGetAllTask() {
	testCases := []struct {
		name  string
		tasks []domain.Task
	}{
		{
			name: "OK - get all tasks",
			tasks: []domain.Task{
				{
					ID:          1,
					Name:        "test1",
					Description: nil,
					Status:      domain.TaskStatusCompleted,
					CreatedAt:   time.Now().UTC(),
					UpdatedAt:   time.Now().UTC(),
				},
				{
					ID:          2,
					Name:        "test2",
					Description: utils.PointerFromValue("description"),
					Status:      domain.TaskStatusActive,
					CreatedAt:   time.Now().UTC(),
					UpdatedAt:   time.Now().UTC(),
				},
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tasks, err := s.storage.GetTasks()

			assert.NoError(s.T(), err)

			for i := range tasks {
				assert.Equal(s.T(), tc.tasks[i].ID, tasks[i].ID)
				assert.Equal(s.T(), tc.tasks[i].Name, tasks[i].Name)
				assert.Equal(s.T(),
					utils.ValueFromPointer(tc.tasks[i].Description),
					utils.ValueFromPointer(tasks[i].Description),
				)
				assert.Equalf(s.T(), tc.tasks[i].Status, tasks[i].Status, fmt.Sprintf("task id: %d", tasks[i].ID))
			}
		})
	}
}

func (s *TaskRepositoryTestSuite) TestUpdateTask() {
	testCases := []struct {
		name          string
		id            int
		status        domain.TaskStatus
		errorExpected bool
	}{
		{
			name:          "OK - task exists",
			id:            2,
			status:        domain.TaskStatusCancelled,
			errorExpected: false,
		},
		{
			name:          "Error - task does not exist",
			id:            100,
			status:        domain.TaskStatusCompleted,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := s.storage.UpdateTaskStatus(tc.id, tc.status)

			if tc.errorExpected {
				assert.Error(s.T(), err)
			} else {
				assert.NoError(s.T(), err)
			}
		})
	}
}

func (s *TaskRepositoryTestSuite) TestRemoveTask() {
	testCases := []struct {
		name          string
		id            int
		errorExpected bool
	}{
		{
			name:          "OK - task exists",
			id:            1,
			errorExpected: false,
		},
		{
			name:          "Error - task does not exist",
			id:            100,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := s.storage.RemoveTask(tc.id)

			if tc.errorExpected {
				assert.Error(s.T(), err)
			} else {
				assert.NoError(s.T(), err)
			}
		})
	}
}

func TestStorage(t *testing.T) {
	const storagePath = "./storage/test_storage.db"

	storage, err := New(Config{
		StoragePath: storagePath,
	})

	assert.NoError(t, err)
	assert.NotNil(t, storage)

	err = storage.Close()

	assert.NoError(t, err)
}

func TestFakeStorage(t *testing.T) {
	fakeStorage := Storage{db: nil}
	err := fakeStorage.Close()

	assert.Error(t, err)
}
