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
	Index         int   `bson:"index" json:"index,omitempty"`
	Points        []int `bson:"points" json:"points,omitempty"`
	RoundResult   int   `bson:"roundResult" json:"round_result,omitempty"`
	InterimResult int   `bson:"interimResult" json:"interim_result,omitempty"`
	Count10       int   `bson:"count10" json:"count_10,omitempty"`
	Count9        int   `bson:"count9" json:"count_9,omitempty"`
}

type CompetitionQualificationEntity struct {
	CompetitionUUID string                           `bson:"competition_uuid" json:"competition_uuid"`
	RM              []QualificationParticipantScores `bson:"rm" json:"rm"`   // recursive men
	RW              []QualificationParticipantScores `bson:"rw" json:"rw"`   // recursive women
	CM              []QualificationParticipantScores `bson:"cm" json:"cm"`   // compound men
	CW              []QualificationParticipantScores `bson:"cw" json:"cw"`   // compound women
	JRM             []QualificationParticipantScores `bson:"jrw" json:"jrw"` // junior recursive men
	JRW             []QualificationParticipantScores `bson:"jrm" json:"jrm"` // junior recursive women
	JCM             []QualificationParticipantScores `bson:"jcw" json:"jcw"` // junior compound men
	JCW             []QualificationParticipantScores `bson:"jcm" json:"jcm"` // junior compound women
}
