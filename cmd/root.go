package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var LineFlag bool

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wc-cli",
		Short: "Print newline, word, and byte counts for each file",
		Long: `wc-cli is a CLI library for Go prints newline, 
		word and byte counts for each file.`,

		Run: CountLines,
	}
	cmd.PersistentFlags().BoolVarP(&LineFlag, "line", "l", false, "print the newline counts")
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

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd().Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
