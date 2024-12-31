package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CountBytes(cmd *cobra.Command, args []string) {

	for _, val := range args {

		data, err := os.ReadFile(val)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "%v", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%d %s\n", len(data), val)
	}
}