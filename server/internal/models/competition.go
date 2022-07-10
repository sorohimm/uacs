package models

import (
	"github.com/google/uuid"
	"time"
)

type Competition struct {
	UUID                   string    `bson:"uuid" json:"uuid" validate:"omitempty"`
	CreatorUUID            string    `bson:"creatorUUID" json:"creatorUUID" validate:"required"`
	Code                   string    `bson:"code" json:"code" validate:"required"`
	Name                   string    `bson:"name" json:"name" validate:"required"`
	Description            string    `bson:"description" json:"description"`
	ShortName              string    `bson:"shortName" json:"shortName"`
	OrganizedBy            string    `bson:"organizedBy" json:"organizedBy" validate:"required"`
	OrganizedByDescription string    `bson:"organizedByDescription" json:"organizedByDescription"`
	CompetitionRules       string    `bson:"competitionRules" json:"competitionRules" validate:"required"`
	TormentType            string    `bson:"tormentType" json:"tormentType"`
	AgeCategories          string    `bson:"ageCategories" json:"ageCategories" validate:"required"`
	Venue                  string    `bson:"venue" json:"venue"`
	Country                string    `bson:"country" json:"country" validate:"omitempty"`
	City                   string    `bson:"city" json:"city" validate:"required"`
	DateFrom               time.Time `bson:"dateFrom" json:"dateFrom" validate:"required"`
	DateTo                 time.Time `bson:"dateTo" json:"dateTo" validate:"required"`
	TimeZone               string    `bson:"timeZone" json:"timeZone"`
	LastUpdate             string    `bson:"lastUpdate" json:"lastUpdate"`
}

func (c *Competition) GenerateUUID() {
	c.UUID = uuid.New().String()
}

type CompetitionShortOutput struct {
	UUID        string `bson:"uuid" json:"uuid"`
	Code        string `bson:"code" json:"code"`
	Name        string `bson:"name" json:"name"`
	OrganizedBy string `bson:"organizedBy" json:"organizedBy"`
	Location    string `bson:"location" json:"location"`
	Date        string `bson:"date" json:"date"`
	LastUpdate  string `bson:"lastUpdate" json:"lastUpdate"`
}

type CompetitionJudgesEntity struct {
	CompetitionUUID string             `bson:"competitionUUID" json:"competitionUUID"`
	JudgingStaff    []CompetitionJudge `bson:"judgingStaff" json:"judgingStaff"`
}

type CompetitionParticipantsEntitySorted struct {
	CompetitionUUID string                                `bson:"competitionUUID" json:"competitionUUID"`
	Compound        CompetitionDivisionParticipantsEntity `bson:"compound" json:"compound"`   // all compound
	Recursive       CompetitionDivisionParticipantsEntity `bson:"recursive" json:"recursive"` // all recursive
}

const MensDivision = "mens"
const WomenDivision = "women"
const Under21MenDivision = "U21M"
const Under21WomenDivision = "U21W"
const Under18MenDivision = "U18M"
const Under18WomenDivision = "U18W"

type CompetitionDivisionParticipantsEntity struct {
	Mens  []CompetitionParticipant `bson:"mens" json:"mens"`   // mens
	Women []CompetitionParticipant `bson:"women" json:"women"` // women
	U21M  []CompetitionParticipant `bson:"U21M" json:"U21M"`   // under 21 men
	U21W  []CompetitionParticipant `bson:"U21W" json:"U21W"`   // under 21 women
	U18M  []CompetitionParticipant `bson:"U18M" json:"U18M"`   // under 18 men
	U18W  []CompetitionParticipant `bson:"U18W" json:"U18W"`   // under 18 women
}

type CompetitionParticipantsEntity struct {
	CompetitionUUID string                   `bson:"competitionUUID" json:"competitionUUID"`
	Participants    []CompetitionParticipant `bson:"participants" json:"participants"` // participants
}

func (e CompetitionParticipantsEntitySorted) ToShortOutput() []CompetitionParticipantShortOutput {
	var output []CompetitionParticipantShortOutput
	for _, el := range e.Compound.Mens {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Compound.Women {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Compound.U21M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Compound.U21M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Compound.U18M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Compound.U21W {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.Mens {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.Women {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.U21M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.U21M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.U18M {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	for _, el := range e.Recursive.U21W {
		participant := CompetitionParticipantShortOutput{
			UUID:        el.UUID,
			FirstName:   el.FirstName,
			LastName:    el.LastName,
			Region:      el.Region,
			RegionCode:  "TODO",
			Division:    el.Division,
			AgeCategory: el.AgeCategory,
		}
		output = append(output, participant)
	}

	return output
}
