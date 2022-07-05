package models

import (
	"github.com/google/uuid"
	"time"
)

type Competition struct {
	UUID                   string    `bson:"uuid" json:"uuid" validate:"omitempty"`
	CreatorUUID            string    `bson:"creatorUUID" json:"creator_uuid" validate:"required"`
	Code                   string    `bson:"code" json:"code" validate:"required"`
	Name                   string    `bson:"name" json:"name" validate:"required"`
	Description            string    `bson:"description" json:"description"`
	ShortName              string    `bson:"shortName" json:"short_name"`
	OrganizedBy            string    `bson:"organizedBy" json:"organized_by" validate:"required"`
	OrganizedByDescription string    `bson:"organizedByDescription" json:"organized_by_description"`
	CompetitionRules       string    `bson:"competitionRules" json:"competition_rules" validate:"required"`
	TormentType            string    `bson:"tormentType" json:"torment_type"`
	AgeCategories          string    `bson:"ageCategories" json:"age_categories" validate:"required"`
	Venue                  string    `bson:"venue" json:"venue"`
	Country                string    `bson:"country" json:"country" validate:"omitempty"`
	City                   string    `bson:"city" json:"city" validate:"required"`
	DateFrom               time.Time `bson:"dateFrom" json:"date_from" validate:"required"`
	DateTo                 time.Time `bson:"dateTo" json:"date_to" validate:"required"`
	TimeZone               string    `bson:"timeZone" json:"time_zone"`
	LastUpdate             string    `bson:"lastUpdate" json:"last_update"`
}

func (c *Competition) GenerateUUID() {
	c.UUID = uuid.New().String()
}

type CompetitionShortOutput struct {
	UUID        string `bson:"uuid" json:"uuid"`
	Code        string `bson:"code" json:"code"`
	Name        string `bson:"name" json:"name"`
	OrganizedBy string `bson:"organizedBy" json:"organized_by"`
	Location    string `bson:"location" json:"location"`
	Date        string `bson:"date" json:"date"`
	LastUpdate  string `bson:"lastUpdate" json:"last_update"`
}

type CompetitionJudgesEntity struct {
	CompetitionUUID string             `bson:"competition_uuid" json:"competition_uuid"`
	JudgingStaff    []CompetitionJudge `bson:"judging_staff" json:"judging_staff"`
}

type CompetitionParticipantsEntity struct {
	CompetitionUUID string                                `bson:"competition_uuid" json:"competition_uuid"`
	Compound        CompetitionDivisionParticipantsEntity `bson:"compound" json:"compound"`   // all compound
	Recursive       CompetitionDivisionParticipantsEntity `bson:"recursive" json:"recursive"` // all recursive
}

type CompetitionDivisionParticipantsEntity struct {
	Mens  []CompetitionParticipant `bson:"mens" json:"mens"`   // mens
	Women []CompetitionParticipant `bson:"women" json:"women"` // women
	U21M  []CompetitionParticipant `bson:"U21M" json:"U21M"`   // under 21 men
	U21W  []CompetitionParticipant `bson:"U21W" json:"U21W"`   // under 21 women
	U18M  []CompetitionParticipant `bson:"U18M" json:"U18M"`   // under 18 men
	U18W  []CompetitionParticipant `bson:"U18W" json:"U18W"`   // under 18 women
}
