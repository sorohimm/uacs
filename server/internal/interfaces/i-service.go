package interfaces

import "uacs/internal/models"

type ICompetitionServices interface {
	CreateCompetition(competition models.Competition) (models.Competition, error)
	GetMyCompetitionsShort(uid string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(id string) (models.Competition, error)
	DeleteCompetition(id string) error
	UpdateCompetition(competition models.Competition) (models.Competition, error)
}

type IParticipantsServices interface {
	GetParticipants(competitionId string) (models.CompetitionParticipantsEntity, error)
	GetParticipant(competitionId string, id string) (models.CompetitionParticipant, error)
	DeleteParticipant(competitionId string, id string) error
	UpdateParticipant(competitionId string, participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
	CreateParticipant(competitionId string, participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
}

type IJudgesServices interface {
	GetJudges(competitionId string) (models.CompetitionJudge, error)
	GetJudge(competitionId string, id string) ([]models.CompetitionJudge, error)
	UpdateJudge(competitionId string, judge models.CompetitionJudge) (models.CompetitionJudge, error)
	DeleteJudge(competitionId string, id string) error
	CreateJudge(competitionId string, judge models.CompetitionJudge) (models.CompetitionJudge, error)
}
