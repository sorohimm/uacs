package models

import "time"

type WarmUpSchedule struct {
	Day       string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Time
	Comment   string
	Targets   string
}

type QualificationRoundSchedule struct {
	Day       string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Time
	Comment   string
	WarmUp    WarmUpSchedule
}
