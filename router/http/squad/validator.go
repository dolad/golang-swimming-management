package squad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
	"swimming-content-management/data/squad"
)

type fieldError struct {
	err validator.FieldError
}

type SquadDataValidator struct {
	Name string `json:"name" binding:"required" `
}
type AddCoachesValidator struct {
	CoachId string `json:"coach_id" binding:"required"`
}

type AddSwimmerValidator struct {
	SwimmerId string `json:"swimmer_id" binding:"required"`
}

func Bind(c *gin.Context) (*squad.Squad, error) {
	var json SquadDataValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}
	squad := &squad.Squad{
		Name: json.Name,
	}
	return squad, nil
}

func BindAddCoach(c *gin.Context) (string, error) {
	var json AddCoachesValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", err
	}
	return json.CoachId, nil
}

func BindAddSwimmer(c *gin.Context) (string, error) {
	var json AddSwimmerValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", err
	}
	return json.SwimmerId, nil
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
