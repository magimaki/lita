package cli

import (
	"fmt"
	"lita"
	"lita/model"
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
	Long:  "Set event reminder.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(remindCmdFlagDate) != 8 {
			fmt.Println(lita.NewError(lita.ErrCmdInput, "invalid event date format, need `yyyymmdd`", nil))
			os.Exit(1)
		}

		if len(remindCmdFlagTime) != 4 && len(remindCmdFlagTime) != 0 {
			fmt.Println(remindCmdFlagTime)
			fmt.Println(len(remindCmdFlagTime))
			fmt.Println(lita.NewError(lita.ErrCmdInput, "invalid event time format, need `hhmm`", nil))
			os.Exit(1)
		}

		var onlyDateEvent bool
		if len(remindCmdFlagTime) == 0 {
			remindCmdFlagTime = "0800"
			onlyDateEvent = true
		}

		year := remindCmdFlagDate[0:4]
		month := remindCmdFlagDate[4:6]
		day := remindCmdFlagDate[6:]
		hour := remindCmdFlagTime[0:2]
		minute := remindCmdFlagTime[2:]
		timeStr := fmt.Sprintf("%s-%s-%s %s:%s", year, month, day, hour, minute)
		timeObj, err := model.StrToTime(timeStr)
		if err != nil {
			fmt.Println(lita.NewError(lita.ErrCmdInput,
				"invalid event date or time input, need `yyyymmdd` (m 01-12, d 01-31) "+
					"and `hhmm` (h 01-23 m 01-59) with only digit", err))
			os.Exit(1)
		}
		reminder := model.Reminder{
			Year:          year,
			Month:         month,
			Day:           day,
			Hour:          hour,
			Minute:        minute,
			Weekday:       timeObj.Weekday().String(),
			TimeStr:       timeStr,
			Time:          timeObj,
			Event:         remindCmdFlagEvent,
			Address:       remindCmdFlagAddress,
			OnlyDateEvent: onlyDateEvent,
		}
		fmt.Println(reminder)
	},
}

func init() {
	remindCmd.Flags().StringVarP(&remindCmdFlagDate,
		"date", "d", "", "event date: yyyymmdd")
	remindCmd.Flags().StringVarP(&remindCmdFlagTime,
		"time", "t", "", "optional event time: hhmm, default 08:00")
	remindCmd.Flags().StringVarP(&remindCmdFlagEvent,
		"event", "e", "", "event content")
	remindCmd.Flags().StringVarP(&remindCmdFlagAddress,
		"address", "a", "", "optional event address")

	if err := remindCmd.MarkFlagRequired("date"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := remindCmd.MarkFlagRequired("event"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	remindCmd.MarkFlagsRequiredTogether("date", "event")
}
