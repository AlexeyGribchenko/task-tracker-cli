package colors

import "testing"

func TestNoopColorer(t *testing.T) {
	n := &NoopColorer{}

	testCases := []struct {
		name   string
		method func(string) string
		input  string
	}{
		{"Bold", n.Bold, "hello"},
		{"Red", n.Red, "world"},
		{"Green", n.Green, "foo"},
		{"Yellow", n.Yellow, "bar"},
		{"Blue", n.Blue, "baz"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.method(tc.input)
			if got != tc.input {
				t.Errorf("%s() = %v, want %v", tc.name, got, tc.input)
			}
		})
	}
}
