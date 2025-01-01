package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CountLines(cmd *cobra.Command, args []string) {

	var lineCount int = 0

	// Scan from standard input if no file is provided
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d\n", lineCount)
	}

	// Else scan from the provided file(s)
	for _, val := range args {
		file, err := os.Open(val)
		if err != nil {
			fmt.Printf("Error occured while opening file: %v", err)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lineCount++
		}
		if err = scanner.Err(); err != nil {
			fmt.Printf("Error occured while reading file: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", lineCount, val)
	}
}
