package usecase

import (
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/dto"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase/mocks"
)

func TestUpdateUseCase(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskRepository, input dto.UpdateTask)

	testCases := []struct {
	}{}

	_ = testCases
}
