package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	LineFlag bool
	ByteFlag bool
)

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wc-cli",
		Short: "Print newline, word, and byte counts for each file",
		Long: `wc-cli is a CLI library for Go prints newline, 
		word and byte counts for each file.`,

		Run: func(cmd *cobra.Command, args []string) {
			if ByteFlag {
				CountBytes(cmd, args)
			} else if LineFlag {
				CountLines(cmd, args)
			} else {
				fmt.Println("No flags passed")
			}
		},
	}
	cmd.PersistentFlags().BoolVarP(&LineFlag, "line", "l", false, "print the newline counts")
	cmd.PersistentFlags().BoolVarP(&ByteFlag, "byte", "m", false, "print number of bytes")

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
