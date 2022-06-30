package models

import "time"

type WarmUpSchedule struct {
	Day       string    `json:"day,omitempty"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  time.Time `json:"duration"`
	Comment   string    `json:"comment,omitempty"`
	Targets   string    `json:"targets,omitempty"`
}

type QualificationRoundSchedule struct {
	Day       string         `json:"day,omitempty"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	Duration  time.Time      `json:"duration"`
	Comment   string         `json:"comment,omitempty"`
	WarmUp    WarmUpSchedule `json:"warm_up"`
}
