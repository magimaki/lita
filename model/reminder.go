package model

import (
	"time"
)

var (
	TimeTemplate = "2006-01-02 15:04"
)

type Reminder struct {
	Year    string `json:"year"`
	Month   string `json:"month"`
	Day     string `json:"day"`
	Hour    string `json:"hour"`
	Minute  string `json:"minute"`
	Weekday string `json:"weekday"`

	// TimeStr has format like `2006-01-02 15:04`.
	TimeStr string `json:"time_str"`
	// Time is a time.Time object contains Year, Month, Day, Hour, Minute.
	Time time.Time `json:"time"`

	Event   string `json:"event"`
	Address string `json:"address"`
}

func StrToTime(timeStr string) (*time.Time, error) {
	timeObj, err := time.Parse(TimeTemplate, timeStr)
	if err != nil {
		return nil, err
	}
	return &timeObj, err
}
