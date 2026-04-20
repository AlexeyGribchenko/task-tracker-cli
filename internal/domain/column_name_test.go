package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseColumnName(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		errorExpected  bool
		expectedError  error
		expectedResult ColumnName
	}{
		{
			name:           "OK - valid name: id",
			input:          "id",
			errorExpected:  false,
			expectedResult: ColumnName(columnId),
		},
		{
			name:           "OK - valid name: name",
			input:          "name",
			errorExpected:  false,
			expectedResult: ColumnName(columnName),
		},
		{
			name:           "OK - valid name: description",
			input:          "description",
			errorExpected:  false,
			expectedResult: ColumnName(columnDescription),
		},
		{
			name:           "OK - valid name: created_at",
			input:          "created_at",
			errorExpected:  false,
			expectedResult: ColumnName(columnCreatedAt),
		},
		{
			name:           "OK - valid name: updated_at",
			input:          "updated_at",
			errorExpected:  false,
			expectedResult: ColumnName(columnUpdatedAt),
		},
		{
			name:           "OK - valid name: status",
			input:          "status",
			errorExpected:  false,
			expectedResult: ColumnName(columnStatus),
		},
		{
			name:           "OK - valid name: upper register",
			input:          "Status",
			errorExpected:  false,
			expectedResult: ColumnName(columnStatus),
		},
		{
			name:          "Error - invalid name: invalid",
			input:         "invalid",
			errorExpected: true,
			expectedError: ErrInvalidColumnName,
		},
		{
			name:          "Error - invalid name: empty",
			input:         "",
			errorExpected: true,
			expectedError: ErrInvalidColumnName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := ParseColumnName(tc.input)

			if tc.errorExpected {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expectedResult)
			}
		})
	}
}
