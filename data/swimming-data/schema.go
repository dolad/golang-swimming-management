package swimming_data

import "github.com/jinzhu/gorm"

type SwimmingData struct {
	gorm.Model
	TotalDistanceCovered uint32
	StrokeCount          uint32
	HeartRate            uint32
	SwimmingType         string
	TimeTakenInSeconds   uint32
	UserID               uint32
}
