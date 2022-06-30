package models

import "time"

type Competition struct {
	UUID                   string    `json:"uuid,omitempty"`
	Code                   string    `json:"code,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Description            string    `json:"description,omitempty"`
	ShortName              string    `json:"short_name,omitempty"`
	OrganizedBy            string    `json:"organized_by,omitempty"`
	OrganizedByDescription string    `json:"organized_by_description,omitempty"`
	CompetitionRules       string    `json:"competition_rules,omitempty"`
	TormentType            string    `json:"torment_type,omitempty"`
	AgeCategories          string    `json:"age_categories,omitempty"`
	Venue                  string    `json:"venue,omitempty"`
	Country                string    `json:"country,omitempty"`
	DateFrom               time.Time `json:"date_from"`
	DateTo                 time.Time `json:"date_to"`
	TimeZone               string    `json:"time_zone,omitempty"`
}
