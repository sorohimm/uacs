package models

type CompetitionParticipant struct {
	UUID              string `json:"uuid"`
	IRMStatus         string `json:"irm_status,omitempty"`
	Code              string `json:"code,omitempty"`
	Shift             string `json:"shift,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	Sex               string `json:"sex,omitempty"`
	BirthDate         string `json:"birth_date,omitempty"`
	TargetNumber      string `json:"target_number,omitempty"`
	Region            string `json:"region,omitempty"`
	RegionName        string `json:"region_name,omitempty"`
	Region2           string `json:"region_2,omitempty"`
	RegionName2       string `json:"region_name_2,omitempty"`
	Region3           string `json:"region_3,omitempty"`
	RegionName3       string `json:"region_name_3,omitempty"`
	Division          string `json:"division,omitempty"`
	Class             string `json:"class,omitempty"`
	AgeClass          string `json:"age_class,omitempty"`
	Discharge         string `json:"discharge,omitempty"`
	IsIndividual      bool   `json:"is_individual,omitempty"`
	IsTeam            bool   `json:"is_team,omitempty"`
	IsIndividualFinal bool   `json:"is_individual_final,omitempty"`
	IsTeamFinal       bool   `json:"is_team_final,omitempty"`
	IsMixFinal        bool   `json:"is_mix_final,omitempty"`
	IsWheelchair      bool   `json:"is_wheelchair,omitempty"`
	Email             string `json:"email,omitempty"`
}

type CompetitionParticipantShortOutput struct {
	Name            string                `json:"name,omitempty"`
	Region          string                `json:"region,omitempty"`
	RegionCode      string                `json:"region_code,omitempty"`
	DistancesPoints []DistancePointsShort `json:"distances_points,omitempty"`
	FinalPoints     int                   `json:"final_points,omitempty"`
	Count10         int                   `json:"count_10,omitempty"`
	Count9          int                   `json:"count_9,omitempty"`
}

type DistancePointsShort struct {
	Index    int `json:"index,omitempty"`
	Distance int `json:"distance,omitempty"`
	Count10  int `json:"count_10,omitempty"`
	Points   int `json:"points,omitempty"`
}

type CompetitionJudge struct {
	UUID       string `json:"uuid"`
	Code       string `json:"code,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Sex        string `json:"sex,omitempty"`
	RegionCode string `json:"region_code,omitempty"`
	Region     string `json:"region,omitempty"`
	Category   string `json:"category,omitempty"`
}
