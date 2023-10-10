package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " Hello ",
			expected: []string{"hello"},
		},
		{
			input:    " Hello WORLD",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		output := cleanInput(c.input)
		if len(output) != len(c.expected) {
			t.Errorf("lengths don't match: %v, want: %v", output, c.expected)
		}
		for i := range output {
			word := output[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words don't match: %v, want: %v", len(output), len(c.expected))
			}
		}
	}
}
