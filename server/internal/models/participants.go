package models

import "github.com/google/uuid"

type CompetitionParticipant struct {
	UUID              string `bson:"uuid" json:"uuid"`
	IRMStatus         string `bson:"IRMStatus" json:"irm_status,omitempty"`
	Code              string `bson:"code" json:"code,omitempty"`
	Shift             string `bson:"shift" json:"shift,omitempty"`
	FirstName         string `bson:"firstName" json:"first_name,omitempty"`
	LastName          string `bson:"lastName" json:"last_name,omitempty"`
	Sex               string `bson:"sex" json:"sex,omitempty"`
	BirthDate         string `bson:"birthDate" json:"birth_date,omitempty"`
	TargetNumber      string `bson:"targetNumber" json:"target_number,omitempty"`
	Region            string `bson:"region" json:"region,omitempty"`
	RegionName        string `bson:"regionName" json:"region_name,omitempty"`
	Region2           string `bson:"region2" json:"region_2,omitempty"`
	RegionName2       string `bson:"regionName2" json:"region_name_2,omitempty"`
	Region3           string `bson:"region3" json:"region_3,omitempty"`
	RegionName3       string `bson:"regionName3" json:"region_name_3,omitempty"`
	Division          string `bson:"division" json:"division,omitempty"`
	AgeCategory       string `bson:"class" json:"class,omitempty"`
	AgeClass          string `bson:"ageClass" json:"age_class,omitempty"`
	Discharge         string `bson:"discharge" json:"discharge,omitempty"`
	IsIndividual      bool   `bson:"isIndividual" json:"is_individual,omitempty"`
	IsTeam            bool   `bson:"isTeam" json:"is_team,omitempty"`
	IsIndividualFinal bool   `bson:"isIndividualFinal" json:"is_individual_final,omitempty"`
	IsTeamFinal       bool   `bson:"isTeamFinal" json:"is_team_final,omitempty"`
	IsMixFinal        bool   `bson:"isMixFinal" json:"is_mix_final,omitempty"`
	IsWheelchair      bool   `bson:"isWheelchair" json:"is_wheelchair,omitempty"`
	Email             string `bson:"email" json:"email,omitempty"`
}

func (c *CompetitionParticipant) GenerateUUID() {
	c.UUID = uuid.New().String()
}

type CompetitionParticipantShortOutput struct {
	UUID            string                `bson:"uuid" json:"uuid"`
	FirstName       string                `bson:"firstName" json:"first_name,omitempty"`
	LastName        string                `bson:"lastName" json:"last_name,omitempty"`
	Division        string                `bson:"division" json:"division"`
	AgeCategory     string                `bson:"ageCategory" json:"age_category"`
	Region          string                `bson:"region" json:"region,omitempty"`
	RegionCode      string                `bson:"regionCode" json:"region_code,omitempty"`
	DistancesPoints []DistancePointsShort `bson:"distancesPoints" json:"distances_points,omitempty"`
	FinalPoints     int                   `bson:"finalPoints" json:"final_points,omitempty"`
	Count10         int                   `bson:"count10" json:"count_10,omitempty"`
	Count9          int                   `bson:"count9" json:"count_9,omitempty"`
}

type DistancePointsShort struct {
	Index    int `bson:"index" json:"index,omitempty"`
	Distance int `bson:"distance" json:"distance,omitempty"`
	Count10  int `bson:"count10" json:"count_10,omitempty"`
	Points   int `bson:"points" json:"points,omitempty"`
}

type CompetitionJudge struct {
	UUID       string `bson:"uuid" json:"uuid"`
	Code       string `bson:"code" json:"code,omitempty"`
	FirstName  string `bson:"firstName" json:"first_name,omitempty"`
	LastName   string `bson:"lastName" json:"last_name,omitempty"`
	Sex        string `bson:"sex" json:"sex,omitempty"`
	RegionCode string `bson:"regionCode" json:"region_code,omitempty"`
	Region     string `bson:"region" json:"region,omitempty"`
	Category   string `bson:"category" json:"category,omitempty"`
}

func (c *CompetitionJudge) GenerateUUID() {
	c.UUID = uuid.New().String()
}
