package models

import "github.com/google/uuid"

type CompetitionParticipant struct {
	UUID              string `bson:"uuid" json:"uuid"`
	IRMStatus         string `bson:"IRMStatus" json:"IRMStatus,omitempty"`
	Code              string `bson:"code" json:"code"`
	Shift             string `bson:"shift" json:"shift"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Sex               string `bson:"sex" json:"sex"`
	BirthDate         string `bson:"birthDate" json:"birthDate"`
	TargetNumber      string `bson:"targetNumber" json:"targetNumber"`
	Region            string `bson:"region" json:"region"`
	RegionName        string `bson:"regionName" json:"regionName"`
	Region2           string `bson:"region2" json:"region2"`
	RegionName2       string `bson:"regionName2" json:"regionName2"`
	Region3           string `bson:"region3" json:"region3"`
	RegionName3       string `bson:"regionName3" json:"regionName3"`
	Division          string `bson:"division" json:"division"`
	AgeCategory       string `bson:"class" json:"class"`
	AgeClass          string `bson:"ageClass" json:"ageClass"`
	Discharge         string `bson:"discharge" json:"discharge"`
	IsIndividual      bool   `bson:"isIndividual" json:"isIndividual"`
	IsTeam            bool   `bson:"isTeam" json:"isTeam"`
	IsIndividualFinal bool   `bson:"isIndividualFinal" json:"isIndividualFinal"`
	IsTeamFinal       bool   `bson:"isTeamFinal" json:"isTeamFinal"`
	IsMixFinal        bool   `bson:"isMixFinal" json:"isMixFinal"`
	IsWheelchair      bool   `bson:"isWheelchair" json:"isWheelchair"`
	Email             string `bson:"email" json:"email"`
}

func (c *CompetitionParticipant) GenerateUUID() {
	c.UUID = uuid.New().String()
}

type CompetitionParticipantShortOutput struct {
	UUID            string                `bson:"uuid" json:"uuid"`
	FirstName       string                `bson:"firstName" json:"firstName"`
	LastName        string                `bson:"lastName" json:"lastName"`
	Division        string                `bson:"division" json:"division"`
	AgeCategory     string                `bson:"ageCategory" json:"ageCategory"`
	Region          string                `bson:"region" json:"region"`
	RegionCode      string                `bson:"regionCode" json:"regionCode"`
	DistancesPoints []DistancePointsShort `bson:"distancesPoints" json:"distancesPoints"`
	FinalPoints     int                   `bson:"finalPoints" json:"finalPoints"`
	Count10         int                   `bson:"count10" json:"count10"`
	Count9          int                   `bson:"count9" json:"count9"`
}

type DistancePointsShort struct {
	Index    int `bson:"index" json:"index,omitempty"`
	Distance int `bson:"distance" json:"distance,omitempty"`
	Count10  int `bson:"count10" json:"count10,omitempty"`
	Points   int `bson:"points" json:"points,omitempty"`
}

type CompetitionJudge struct {
	UUID       string `bson:"uuid" json:"uuid"`
	Code       string `bson:"code" json:"code"`
	FirstName  string `bson:"firstName" json:"firstName"`
	LastName   string `bson:"lastName" json:"lastName"`
	Sex        string `bson:"sex" json:"sex"`
	RegionCode string `bson:"regionCode" json:"regionCode"`
	Region     string `bson:"region" json:"region"`
	Category   string `bson:"category" json:"category"`
}

func (c *CompetitionJudge) GenerateUUID() {
	c.UUID = uuid.New().String()
}
