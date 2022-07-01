package models

import "time"

type Competition struct {
	UUID                   string    `json:"uuid" validate:"required"`
	CreatorUUID            string    `json:"creator_uuid" validate:"required"`
	Code                   string    `json:"code" validate:"required"`
	Name                   string    `json:"name" validate:"required"`
	Description            string    `json:"description,omitempty"`
	ShortName              string    `json:"short_name,omitempty"`
	OrganizedBy            string    `json:"organized_by" validate:"required"`
	OrganizedByDescription string    `json:"organized_by_description,omitempty"`
	CompetitionRules       string    `json:"competition_rules" validate:"required"`
	TormentType            string    `json:"torment_type,omitempty"`
	AgeCategories          string    `json:"age_categories" validate:"required"`
	Venue                  string    `json:"venue,omitempty"`
	Country                string    `json:"country" validate:"required"`
	DateFrom               time.Time `json:"date_from" validate:"required"`
	DateTo                 time.Time `json:"date_to" validate:"required"`
	TimeZone               string    `json:"time_zone,omitempty"`
	LastUpdate             string    `json:"last_update"`
}

type CompetitionShortOutput struct {
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	OrganizedBy string `json:"organized_by,omitempty"`
	Location    string `json:"location,omitempty"`
	Date        string `json:"date,omitempty"`
	LastUpdate  string `json:"last_update,omitempty"`
}
