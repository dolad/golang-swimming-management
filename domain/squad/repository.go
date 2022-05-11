package squad

import (
	uuid "github.com/satori/go.uuid"
	"swimming-content-management/data/squad"
)

type SquadRepository interface {
	CreateSquad(squadData *squad.Squad) (*squad.Squad, error)
	GetSquads() ([]squad.Squad, error)
	GetSquad(squadId uint32) (squad.Squad, error)
	AddCoachToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error)
	AddSwimmerToSquad(squadId uint32, coachId uuid.UUID) (*squad.Squad, error)
	//GetSquadByCoachName(squadId uuid.UUID) (squad.Squad, error)
}
