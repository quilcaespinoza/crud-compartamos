package utils

import (
	"crud-compartamos/models"
	"time"
)

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}
func (v *Formatter) FormatBirthdate(birthdate string) string {
	parsedDate, err := time.Parse(time.RFC3339, birthdate)
	if err != nil {
		return birthdate
	}

	formattedDate := parsedDate.Format("2006-01-02")
	return formattedDate
}

func (v *Formatter) FormatUsersResult(users interface{}) interface{} {
	if users == nil || len(users.([]models.User)) == 0 {
		return map[string]interface{}{"users": map[string]interface{}{}}
	}

	return map[string]interface{}{"users": users}
}
