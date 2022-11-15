package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	remindCmdFlagDate    string
	remindCmdFlagTime    string
	remindCmdFlagEvent   string
	remindCmdFlagAddress string
)

var remindCmd = &cobra.Command{
	Use:   "remind",
	Short: "Set event reminder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(remindCmdFlagDate)
		fmt.Println(remindCmdFlagTime)
		fmt.Println(remindCmdFlagEvent)
		fmt.Println(remindCmdFlagAddress)
	},
}

func init() {
	remindCmd.Flags().StringVarP(&remindCmdFlagDate, "date", "d", "", "event date: yyyymmdd")
	remindCmd.Flags().StringVarP(&remindCmdFlagTime, "time", "t", "0800", "event time: hhmm")
	remindCmd.Flags().StringVarP(&remindCmdFlagEvent, "event", "e", "", "Source directory to read from")
	remindCmd.Flags().StringVarP(&remindCmdFlagAddress, "address", "a", "", "Source directory to read from")

	err := remindCmd.MarkFlagRequired("date")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = remindCmd.MarkFlagRequired("event")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	remindCmd.MarkFlagsRequiredTogether("date", "event")
}
