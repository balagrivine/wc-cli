package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CountLines(cmd *cobra.Command, args []string) {

	if LineFlag {
		for _, file := range args {
			lines := func(fileName string) int {
				file, err := os.Open(fileName)
				if err != nil {
					fmt.Printf("Error occured while opening file: %v", err)
				}

				scanner := bufio.NewScanner(file)
				lineCount := 0

				for scanner.Scan() {
					lineCount++
				}
				if err = scanner.Err(); err != nil {
					fmt.Printf("Error occured while reading file: %v", err)
				}
				return lineCount
			}(file)

			fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", lines, file)
		}
	}
}
