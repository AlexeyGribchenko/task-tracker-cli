package colors

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		name           string
		input          Config
		expectedResult bool
	}{
		{
			name:           "OK - color = true",
			input:          Config{ColoredOutput: true},
			expectedResult: true,
		},
		{
			name:           "OK - color = false",
			input:          Config{ColoredOutput: false},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			originalValue := color.NoColor

			t.Cleanup(func() {
				color.NoColor = originalValue
			})

			Init(tc.input)

			assert.Equal(t, !color.NoColor, tc.expectedResult)
		})
	}
}
