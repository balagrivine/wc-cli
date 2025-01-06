package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func countWords(cmd *cobra.Command, args []string) error {

	var wordCount int
	// Scan from stdin if no file is provided
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wordCount++
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d\n", wordCount)
	}

	// Else scan from the file(s) provided
	for _, val := range args {
		file, err := os.Open(val)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wordCount++
		}

		if err = scanner.Err(); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", wordCount, val)

		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}
