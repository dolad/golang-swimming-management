package swimmingdata

import (
	uuid "github.com/satori/go.uuid"
	swimming_data "swimming-content-management/data/swimming-data"
	users "swimming-content-management/data/user"
)

type SwimmingDataService interface {
	AddSwimmingDataToUser(swimmingData *swimming_data.SwimmingData) (*users.User, error)
	GetUsersSwimmingData(usersId uuid.UUID) (*[]swimming_data.SwimmingData, error)
	GetUserSwimmingData() (*[]swimming_data.SwimmingData, error)
}

type Service struct {
	repository SwimmingDataRepository
}

func (svc *Service) AddSwimmingDataToUser(swimmingData *swimming_data.SwimmingData) (*users.User, error) {
	return svc.repository.AddSwimmingDataToUser(swimmingData)
}

func (svc *Service) GetUsersSwimmingData(usersId uuid.UUID) (*[]swimming_data.SwimmingData, error) {
	return svc.repository.GetUsersSwimmingData(usersId)
}

func (svc *Service) GetUserSwimmingData() (*[]swimming_data.SwimmingData, error) {
	return svc.repository.GetUserSwimmingData()
}
func NewService(repository SwimmingDataRepository) *Service {
	return &Service{repository: repository}
}
