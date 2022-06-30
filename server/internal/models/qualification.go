package models

type QualificationParticipantScores struct {
	UUID       string                        `json:"uuid,omitempty"`
	Name       string                        `json:"name,omitempty"`
	Code       string                        `json:"code,omitempty"`
	Region     string                        `json:"region,omitempty"`
	RegionCode string                        `json:"region_code,omitempty"`
	Scores     []QualificationDistanceScores `json:"scores,omitempty"`
}

type QualificationDistanceScores struct {
	Index          int                        `json:"index,omitempty"`
	Distance       string                     `json:"distance,omitempty"`
	Rounds         []QualificationRoundScores `json:"rounds,omitempty"`
	DistanceResult int                        `json:"distance_result,omitempty"`
	Count10        int                        `json:"count_10,omitempty"`
	Count9         int                        `json:"count_9,omitempty"`
}

type QualificationRoundScores struct {
	Index         int   `json:"index,omitempty"`
	Points        []int `json:"points,omitempty"`
	RoundResult   int   `json:"round_result,omitempty"`
	InterimResult int   `json:"interim_result,omitempty"`
	Count10       int   `json:"count_10,omitempty"`
	Count9        int   `json:"count_9,omitempty"`
}