package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Empty file", []string{"-w", "test_data/0_line.txt"}, "0 test_data/0_line.txt\n"},
		{"Multiple line file", []string{"-w", "test_data/10_lines.txt"}, "10 test_data/10_lines.txt\n"},
		{"Multiple files", []string{"-w", "test_data/0_line.txt", "test_data/10_lines.txt"},
			"0 test_data/0_line.txt\n10 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBufferString("")
			cmd := RootCmd()

			cmd.SetOut(buf)
			cmd.SetArgs(test.args)
			cmd.Execute()

			out, err := io.ReadAll(buf)
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
