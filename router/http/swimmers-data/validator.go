package swimmerdata

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
	swimming_data "swimming-content-management/data/swimming-data"
)

type fieldError struct {
	err validator.FieldError
}

type SwimmerDataValidator struct {
	TotalDistanceCovered uint32 `json:"total_distance_covered" binding:"required"`
	StrokeCount          uint32 `json:"stroke_count"`
	HeartRate            uint32 `json:"heart_rate" binding:"required"`
	TimeTakenInSeconds   uint32 `json:"time_taken_in_seconds" binding:"required"`
	SwimmingType         string `json:"swimming_type" binding:"required" `
}

func Bind(c *gin.Context) (*swimming_data.SwimmingData, error) {
	var json SwimmerDataValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}
	newUser := &swimming_data.SwimmingData{
		TotalDistanceCovered: json.TotalDistanceCovered,
		StrokeCount:          json.StrokeCount,
		HeartRate:            json.HeartRate,
		TimeTakenInSeconds:   json.TimeTakenInSeconds,
		SwimmingType:         json.SwimmingType,
	}

	return newUser, nil
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
