package models

type QualificationParticipantScores struct {
	ParticipantUUID string                        `bson:"participant_uuid" json:"uuid,omitempty"`
	Scores          []QualificationDistanceScores `bson:"scores" json:"scores,omitempty"`
}

type QualificationDistanceScores struct {
	Index          int                        `bson:"index" json:"index,omitempty"`
	Distance       string                     `bson:"distance" json:"distance,omitempty"`
	Rounds         []QualificationRoundScores `bson:"rounds" json:"rounds,omitempty"`
	DistanceResult int                        `bson:"distanceResult" json:"distance_result,omitempty"`
	Count10        int                        `bson:"count10" json:"count_10,omitempty"`
	Count9         int                        `bson:"count9" json:"count_9,omitempty"`
}

type QualificationRoundScores struct {
	Index         int   `bson:"index" json:"index"`
	Points        []int `bson:"points" json:"points"`
	RoundResult   int   `bson:"roundResult" json:"roundResult"`
	InterimResult int   `bson:"interimResult" json:"interimResult"`
	Count10       int   `bson:"count10" json:"count10"`
	Count9        int   `bson:"count9" json:"count9"`
}

type CompetitionQualificationEntity struct {
	CompetitionUUID string                                      `bson:"competitionUUID" json:"competitionUUID"`
	Compound        CompetitionDivisionParticipantsScoresEntity `bson:"compound" json:"compound"`   // all compound
	Recursive       CompetitionDivisionParticipantsScoresEntity `bson:"recursive" json:"recursive"` // all recursive
}

type CompetitionDivisionParticipantsScoresEntity struct {
	Mens  []QualificationParticipantScores `bson:"mens" json:"mens"`   // mens
	Women []QualificationParticipantScores `bson:"women" json:"women"` // women
	U21M  []QualificationParticipantScores `bson:"U21M" json:"U21M"`   // under 21 men
	U21W  []QualificationParticipantScores `bson:"U21W" json:"U21W"`   // under 21 women
	U18M  []QualificationParticipantScores `bson:"U18M" json:"U18M"`   // under 18 men
	U18W  []QualificationParticipantScores `bson:"U18W" json:"U18W"`   // under 18 women
}
