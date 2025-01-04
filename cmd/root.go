package cmd

import (
	"os"
	"errors"

	"github.com/spf13/cobra"
)

var (
	lineFlag bool
	byteFlag bool
	wordFlag bool
)

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wc-cli",
		Short: "Print newline, word, and byte counts for each file",
		Long: `wc-cli is a CLI library for Go prints newline, 
		word and byte counts for each file.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if byteFlag {
				return countBytes(args)
			}
			if lineFlag {
				return countLines(args)
			} 
			if wordFlag {
				return countWords(args)
			}
			return errors.New("No flags passed")
			
		},
	}
	cmd.PersistentFlags().BoolVarP(&lineFlag, "line", "l", false, "print the newline counts")
	cmd.PersistentFlags().BoolVarP(&byteFlag, "byte", "m", false, "print number of bytes")
	cmd.PersistentFlags().BoolVarP(&wordFlag, "word", "w", false, "print the word count")

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := RootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
