package models

import "time"

type WarmUpSchedule struct {
	Day       string    `bson:"day" json:"day,omitempty"`
	StartTime time.Time `bson:"startTime" json:"startTime"`
	EndTime   time.Time `bson:"endTime" json:"endTime"`
	Duration  time.Time `bson:"duration" json:"duration"`
	Comment   string    `bson:"comment" json:"comment"`
	Targets   string    `bson:"targets" json:"targets"`
}

type QualificationRoundSchedule struct {
	Day       string         `bson:"day" json:"day,omitempty"`
	StartTime time.Time      `bson:"startTime" json:"startTime"`
	EndTime   time.Time      `bson:"endTime" json:"endTime"`
	Duration  time.Time      `bson:"duration" json:"duration"`
	Comment   string         `bson:"comment" json:"comment"`
	WarmUp    WarmUpSchedule `bson:"warmUp" json:"warmUp"`
}
