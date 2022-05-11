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
	db.AutoMigrate(&users.User{}, &SwimmingData{})
	return &Store{
		db: db,
	}
}

func (s *Store) AddSwimmingDataToUser(swimmingData *SwimmingData) (*users.User, error) {
	//check if user is a swimmer
	result := &users.User{}

	swimmingDataEntity := SwimmingData{
		TotalDistanceCovered: swimmingData.TotalDistanceCovered,
		StrokeCount:          swimmingData.StrokeCount,
		HeartRate:            swimmingData.HeartRate,
		TimeTakenInSeconds:   swimmingData.TimeTakenInSeconds,
		SwimmingType:         swimmingData.SwimmingType,
		UserID:               swimmingData.UserID,
	}
	if err := s.db.Create(&swimmingDataEntity).Error; err != nil {
		return nil, err
	}

	updatedUser := s.db.Preload("SwimmingData").First(&result, "id =?", swimmingData.UserID).Error
	if updatedUser != nil {
		return nil, errors.New("Cant get updated users")
	}
	return result, nil
}

func (s *Store) GetUsersSwimmingData(usersId uuid.UUID) (*[]SwimmingData, error) {
	var result []SwimmingData
	query := s.db.Where("user_id =?", usersId).Find(&result)
	if query.Error != nil {
		return nil, errors.New("Cannot create find users with this id")
	}
	return &result, nil
}

func (s *Store) GetUserSwimmingData() (*[]SwimmingData, error) {
	var result []SwimmingData
	query := s.db.Preload("User").Find(&result)
	if query.Error != nil {
		return nil, errors.New("Cannot create find users with this id")
	}
	return &result, nil
}
