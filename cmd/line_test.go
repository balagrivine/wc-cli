package cmd

import (
	"bytes"
	"testing"
	"io"
)

func TestLine(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected string
	}{
		{"Empty file", "test_data/0_line.txt", "0 test_data/0_line.txt\n"},
		{"Multiple lines", "test_data/10_lines.txt", "10 test_data/10_lines.txt\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := RootCmd()
			var b = bytes.NewBufferString("")

			cmd.SetOut(b)
			cmd.SetArgs([]string{"-l", test.fileName})
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
