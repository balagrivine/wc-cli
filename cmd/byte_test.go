package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestByteCount(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Empty file", []string{"test_data/0_line.txt"}, "0 test_data/0_line.txt\n"},
		{"Multiple line file", []string{"test_data/10_lines.txt"}, "21 test_data/10_lines.txt\n"},
		{"Multiple files", []string{"test_data/0_line.txt", "test_data/10_lines.txt"},
			"0 test_data/0_line.txt\n21 test_data/10_lines.txt\n"},
		{"Non-empty multiple files", []string{"test_data/test.txt", "test_data/10_lines.txt"},
			"20 test_data/test.txt\n21 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bf bytes.Buffer
			originalStdout := os.Stdout

			os.Stdout = &bf

			t.Cleanup(func() {
				os.Stdout = originalStdout
			})

			_ = countBytes(test.args)
			if bf.String() != test.expected {
				t.Errorf("expected %v, got %v\n", test.expected, bf.String())
			}
		},
		)
	}
}
