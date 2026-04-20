package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidTaskStatus = errors.New("invalid task status!")
	ErrInvalidTaskName   = errors.New("invalid task name!")
)

type TaskStatus string

func (t TaskStatus) String() string {
	return string(t)
}

const (
	StrCreated    = "created"
	StrCompleted  = "completed"
	StrInProgress = "in_progress"
	StrCancelled  = "cancelled"

	TaskStatusCreated    TaskStatus = StrCreated
	TaskStatusInProgress TaskStatus = StrInProgress
	TaskStatusCompleted  TaskStatus = StrCompleted
	TaskStatusCancelled  TaskStatus = StrCancelled
)

type Task struct {
	ID          int // Maybe new type TaskID?
	Name        string
	Status      TaskStatus
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(name string, description *string) (*Task, error) {

	if name == "" {
		return nil, ErrInvalidTaskName
	}

	return &Task{
		Name:        name,
		Status:      TaskStatusCreated,
		Description: description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func ParseStatus(input string) (TaskStatus, error) {
	normalized := strings.ToLower(strings.TrimSpace(input))

	switch normalized {
	case StrCreated:
		return TaskStatusCreated, nil
	case StrCancelled:
		return TaskStatusCancelled, nil
	case StrCompleted:
		return TaskStatusCompleted, nil
	case StrInProgress:
		return TaskStatusInProgress, nil
	default:
		return "", ErrInvalidTaskStatus
	}
}
