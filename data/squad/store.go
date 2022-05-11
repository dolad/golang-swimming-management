package squad

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	users "swimming-content-management/data/user"
)

const (
	createError      = "Error in creating new Squad Data for users"
	GettingSquadData = "Error in getting squad data"
	AccessDenied     = "Unauthorized"
	Swimmer          = "swimmer"
	Coach            = "coaches"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	// migrate schema
	db.AutoMigrate(&Squad{})
	return &Store{
		db: db,
	}
}

func (s *Store) CreateSquad(squadData *Squad) (*Squad, error) {
	//check if user is a squad does not exist

	if err := s.db.Omit("Coach ", "Swimmers").Create(&squadData).Error; err != nil {
		return nil, err
	}
	return squadData, nil
}

func (s *Store) GetSquads() ([]Squad, error) {
	var result []Squad
	query := s.db.Preload("Coach").Preload("Swimmers").Find(&result)
	if query.Error != nil {
		return nil, errors.New("Cannot create find users with this id")
	}
	return result, nil
}

func (s *Store) GetSquad(squadId uint32) (Squad, error) {
	var result Squad
	query := s.db.Preload("Coach").Preload("Swimmers").Where("id =?", squadId).Find(&result)

	if query.Error != nil {
		return Squad{}, errors.New("Cannot create find users with this id")
	}
	return result, nil
}

func (s *Store) GetSquadByName(name string) (Squad, error) {
	var result Squad
	query := s.db.Preload("Coach").Preload("Swimmers").Where("name= ?", name).First(&result)

	if query.Error != nil {
		return Squad{}, errors.New("Internal error")
	}

	return result, nil
}

func (s *Store) AddCoachToSquad(squadId uint32, coachId uuid.UUID) (*Squad, error) {
	var result Squad
	var coach users.User
	query := s.db.Preload("Role").Where("id= ?", coachId).First(&coach)
	if query.Error != nil {
		return nil, errors.New("Internal error")
	}
	fmt.Println(coach.Role.Name)
	if coach.Role.Name != Coach {
		return nil, errors.New("This user is not an coach")
	}
	coach.SquadID = squadId
	s.db.Save(&coach)

	updatedSquad := s.db.Preload("Coach").Where("id= ?", squadId).First(&result).Error
	if updatedSquad != nil {
		return nil, errors.New("Cant get updated users")
	}

	return &result, nil
}

func (s *Store) AddSwimmerToSquad(squadId uint32, swimmerId uuid.UUID) (*Squad, error) {
	var result Squad
	var swimmer users.User
	query := s.db.Preload("Role").Where("id= ?", swimmerId).First(&swimmer)
	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, errors.New("Internal error")
	}
	if swimmer.Role.Name != Swimmer {
		return nil, errors.New("This user is not an swimmer")
	}
	swimmer.SquadID = squadId
	s.db.Save(&swimmer)
	updatedSquad := s.db.Preload("Swimmers").Where("id= ?", squadId).First(&result).Error
	if updatedSquad != nil {
		return nil, errors.New("Cant get updated users")
	}

	return &result, nil
}
