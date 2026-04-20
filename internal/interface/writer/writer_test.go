package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsColumnNameValid(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "OK - valid name: id",
			input:          "id",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: name",
			input:          "name",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: description",
			input:          "description",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: created_at",
			input:          "created",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: updated_at",
			input:          "updated",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: status",
			input:          "status",
			expectedResult: true,
		},
		{
			name:           "OK - valid name: upper register",
			input:          "Status",
			expectedResult: true,
		},
		{
			name:           "Error - invalid name: invalid",
			input:          "invalid",
			expectedResult: false,
		},
		{
			name:           "Error - invalid name: empty",
			input:          "",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := isColumnNameValid(tc.input)

			assert.Equal(t, result, tc.expectedResult)
		})
	}
}
