package model

import (
	"time"
)

const (
	TimeTemplate = "2006-01-02 15:04"
)

type Event struct {
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

	Content string `json:"content"`
	Address string `json:"address"`
}

func StrToTime(timeStr string) (time.Time, error) {
	timeObj, err := time.Parse(TimeTemplate, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return timeObj, err
}
