package swimmingdata

import (
	uuid "github.com/satori/go.uuid"
	swimming_data "swimming-content-management/data/swimming-data"
	users "swimming-content-management/data/user"
)

type SwimmingDataRepository interface {
	AddSwimmingDataToUser(swimmingData *swimming_data.SwimmingData) (*users.User, error)
	GetUsersSwimmingData(usersId uuid.UUID) (*[]swimming_data.SwimmingData, error)
	GetUserSwimmingData() (*[]swimming_data.SwimmingData, error)
}
