package swimming_data

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	users "swimming-content-management/data/user"
)

const (
	createError         = "Error in creating new Swimming Data for users"
	GettingSwimmingData = "Error in getting swimming data"
	AccessDenied        = "Password does not match"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	// migrate schema
	return &Store{
		db: db,
	}
}

func (s *Store) AddSwimmingDataToUser(swimmingData *SwimmingData) (*users.User, error) {
	//check if user is a swimmer
	var authUsers users.User
	query := s.db.Preloads("Role").First(authUsers, swimmingData.UserID)

	if query.Error != nil {
		return nil, errors.New("Users not found")
	}
	if authUsers.Role.Name != "swimmer" {
		return nil, errors.New("The current users is not registered as a swimmer")
	}
	swimmingDataEntity := []SwimmingData{
		{TotalDistanceCovered: swimmingData.TotalDistanceCovered,
			StrokeCount:        swimmingData.StrokeCount,
			HeartRate:          swimmingData.HeartRate,
			TimeTakenInSeconds: swimmingData.TimeTakenInSeconds,
		},
	}
	addSwimmingDataQuery := s.db.Model(&authUsers).Association("SwimmingData").Append(swimmingDataEntity)
	if addSwimmingDataQuery.Error != nil {
		return nil, errors.New("Cannot create swimmingdata")
	}
	return &authUsers, nil
}

func (s *Store) GetUsersSwimmingData(usersId uuid.UUID) (*[]SwimmingData, error) {
	var result []SwimmingData
	var authUsers users.User
	query := s.db.Model(&authUsers).Where("user_id", usersId).Association("SwimmingData").Find(&result)
	if query.Error != nil {
		return nil, errors.New("Cannot create find users with this id")
	}
	return &result, nil
}
