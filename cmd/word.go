package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func CountWords(cmd *cobra.Command, args []string) {

	// Scan from stdin if no file is provided
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		wordCount := 0

		for scanner.Scan() {
			wordCount += len(strings.Fields(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "%v\n", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d\n", wordCount)
	}
	
	// Else scan from the file(s) provided
	for _, val := range args {
		file, err := os.Open(val)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "%v\n", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		wordCount := 0

		for scanner.Scan() {
			wordCount += len(strings.Fields(scanner.Text()))
		}

		if err = scanner.Err(); err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "%v\n", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", wordCount, val)
	}
}
