package model

type Reminder struct {
	Evt Event

	// OnlyDateReminder is a reminder without specific starting hour and minute.
	OnlyDateReminder bool `json:"only_date_reminder"`
}
