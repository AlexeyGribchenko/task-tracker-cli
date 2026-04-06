package colors

import "testing"

func TestRealColorer(t *testing.T) {
	r := &RealColorer{}

	testCases := []struct {
		name           string
		method         func(string) string
		expectedPrefix string
		msg            string
	}{
		{"Bold", r.Bold, bold, "bold"},
		{"Red", r.Red, red, "red"},
		{"Green", r.Green, green, "green"},
		{"Yellow", r.Yellow, yellow, "yellow"},
		{"Blue", r.Blue, blue, "blue"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.method(tc.msg)
			expected := tc.expectedPrefix + tc.msg + reset

			if got != expected {
				t.Errorf("%s() = %q, want %q", tc.name, got, expected)
			}
		})
	}
}
