package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(args []string) error {

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

		fmt.Fprintf(os.Stdout, "%d\n", lineCount)
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

		fmt.Fprintf(os.Stdout, "%d %s\n", lineCount, val)
		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}
