package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func countBytes(cmd *cobra.Command, args []string) error {

	// Read from stdin stream if no file is provided
	if len(args) == 0 {
		bytesRead, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d\n", len(bytesRead))
	}

	// Else read from file(s) provided
	for _, val := range args {

		data, err := os.ReadFile(val)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", len(data), val)
	}
	return nil
}
