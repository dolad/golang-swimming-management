package squad

import (
	uuid "github.com/satori/go.uuid"
	"swimming-content-management/data/squad"
)

type SquadDataService interface {
	CreateSquad(squadData *squad.Squad) (*squad.Squad, error)
	GetSquads() ([]squad.Squad, error)
	GetSquad(squadId uint32) (squad.Squad, error)
	AddCoachToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error)
	AddSwimmerToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error)
}

type Service struct {
	repository SquadRepository
}

func (svc *Service) CreateSquad(squadData *squad.Squad) (*squad.Squad, error) {
	return svc.repository.CreateSquad(squadData)
}

func (svc *Service) GetSquads() ([]squad.Squad, error) {
	return svc.repository.GetSquads()
}

func (svc *Service) GetSquad(squadId uint32) (squad.Squad, error) {
	return svc.repository.GetSquad(squadId)
}

func (svc *Service) AddCoachToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error) {
	return svc.repository.AddCoachToSquad(squadId, coachId)
}

func (svc *Service) AddSwimmerToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error) {
	return svc.repository.AddSwimmerToSquad(squadId, coachId)
}

func NewService(repository SquadRepository) *Service {
	return &Service{repository: repository}
}
