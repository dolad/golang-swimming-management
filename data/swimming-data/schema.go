package swimming_data

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	users "swimming-content-management/data/user"
)

type SwimmingData struct {
	gorm.Model
	TotalDistanceCovered uint32
	StrokeCount          uint32
	HeartRate            uint32
	SwimmingType         string
	TimeTakenInSeconds   uint32
	UserID               uuid.UUID
	User                 *users.User
}
