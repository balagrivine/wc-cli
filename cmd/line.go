package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func countLines(cmd *cobra.Command, args []string) error {

	var lineCount uint

	// Scan from standard input if no file is provided
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d", lineCount)
	}

	// Else scan from the provided file(s)
	for _, val := range args {
		file, err := os.Open(val)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lineCount++
		}

		if err = scanner.Err(); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", lineCount, val)
		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}
