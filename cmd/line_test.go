package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestLine(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Empty file", []string{"-l", "test_data/0_line.txt"}, "0 test_data/0_line.txt\n"},
		{"Multiple line file", []string{"-l", "test_data/10_lines.txt"}, "10 test_data/10_lines.txt\n"},
		{"Multiple files", []string{"-l", "test_data/0_line.txt", "test_data/10_lines.txt"},
			"0 test_data/0_line.txt\n10 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := RootCmd()
			var b = bytes.NewBufferString("")

			cmd.SetOut(b)
			cmd.SetArgs(test.args)
			cmd.Execute()

			out, err := io.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}

			if string(out) != test.expected {
				t.Errorf("expected %v, got %v\n", test.expected, string(out))
			}
		},
		)
	}
}
