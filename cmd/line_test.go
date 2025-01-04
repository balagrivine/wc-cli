package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestLineCount(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Empty file", []string{"test_data/0_line.txt"}, "0 test_data/0_line.txt\n"},
		{"Multiple line file", []string{"test_data/10_lines.txt"}, "10 test_data/10_lines.txt\n"},
		{"Multiple files", []string{"test_data/0_line.txt", "test_data/10_lines.txt"},
			"0 test_data/0_line.txt\n10 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bf bytes.Buffer
			originalStdout := os.Stdout
			// Redirect stdout to the buffer for easier access
			os.Stdout = &bf

			t.Cleanup(func() {
				// Restore stdout
				os.Stdout = originalStdout
			})

			_ = countLines(test.args)

			if bf.String() != test.expected {
				t.Errorf("expected %v, got %v\n", test.expected, bf.String())
			}
		},
		)
	}
}
