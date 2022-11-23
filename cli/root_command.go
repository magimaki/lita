package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "lita",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(remindCmd)
	rootCmd.AddCommand(serveCmd)
}

func isValidCmdDateLength(date string) bool {
	return len(date) == 8
}

func isValidCmdTimeLength(time string) bool {
	return len(time) == 4 || len(time) == 0
}
