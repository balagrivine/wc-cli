package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestByteCount(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Empty file", []string{"-m", "test_data/0_line.txt"}, "0 test_data/0_line.txt\n"},
		{"Multiple line file", []string{"-m", "test_data/10_lines.txt"}, "21 test_data/10_lines.txt\n"},
		{"Multiple files", []string{"-m", "test_data/0_line.txt", "test_data/10_lines.txt"},
			"0 test_data/0_line.txt\n21 test_data/10_lines.txt\n"},
		{"Non-empty multiple files", []string{"-m", "test_data/test.txt", "test_data/10_lines.txt"},
			"20 test_data/test.txt\n21 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := RootCmd()
			buf := bytes.NewBufferString("")

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
