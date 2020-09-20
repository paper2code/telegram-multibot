package cmd

import "time"

const (
	PeroidOnce = iota
	PeriodYearly
	PeriodMonthly
	PeriodWeekly
	PeriodDaily
)

var remindPeriods []string = []string{"Once", "Annually", "Monthly", "Weekly", "Daily"}

// Remind is a remind struct
type Remind struct {
	Period int
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
	Time   time.Time
	Name   string
	Text   string
}
