package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueFromPointer(t *testing.T) {
	testCases := []struct {
		name           string
		input          *string
		expectedResult string
	}{
		{
			name:           "OK - value",
			input:          PointerFromValue("test"),
			expectedResult: "test",
		},
		{
			name:           "OK - nil pointer",
			input:          nil,
			expectedResult: "",
		},
		{
			name:           "OK - empty string",
			input:          PointerFromValue(""),
			expectedResult: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValueFromPointer(tc.input)

			assert.Equal(t, result, tc.expectedResult)
		})
	}
}

func TestPointerFromValue(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult *string
	}{
		{
			name:           "OK - value",
			input:          "test",
			expectedResult: PointerFromValue("test"),
		},
		{
			name:           "OK - empty string",
			input:          "",
			expectedResult: PointerFromValue(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := PointerFromValue(tc.input)

			assert.NotNil(t, result)
			assert.Equal(t, *result, *tc.expectedResult)
		})
	}
}
