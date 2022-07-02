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
	Description            string    `bson:"description" json:"description,omitempty"`
	ShortName              string    `bson:"shortName" json:"short_name,omitempty"`
	OrganizedBy            string    `bson:"organizedBy" json:"organized_by" validate:"required"`
	OrganizedByDescription string    `bson:"organizedByDescription" json:"organized_by_description,omitempty"`
	CompetitionRules       string    `bson:"competitionRules" json:"competition_rules" validate:"required"`
	TormentType            string    `bson:"tormentType" json:"torment_type,omitempty"`
	AgeCategories          string    `bson:"ageCategories" json:"age_categories" validate:"required"`
	Venue                  string    `bson:"venue" json:"venue,omitempty"`
	Country                string    `bson:"country" json:"country" validate:"omitempty"`
	City                   string    `bson:"city" json:"city" validate:"required"`
	DateFrom               time.Time `bson:"dateFrom" json:"date_from" validate:"required"`
	DateTo                 time.Time `bson:"dateTo" json:"date_to" validate:"required"`
	TimeZone               string    `bson:"timeZone" json:"time_zone,omitempty"`
	LastUpdate             string    `bson:"lastUpdate" json:"last_update"`
}

func (c *Competition) GenerateUUID() {
	c.UUID = uuid.New().String()
}

type CompetitionShortOutput struct {
	Code        string `bson:"code" json:"code,omitempty"`
	Name        string `bson:"name" json:"name,omitempty"`
	OrganizedBy string `bson:"organizedBy" json:"organized_by,omitempty"`
	Location    string `bson:"location" json:"location,omitempty"`
	Date        string `bson:"date" json:"date,omitempty"`
	LastUpdate  string `bson:"lastUpdate" json:"last_update,omitempty"`
}
