package utils

import (
	"regexp"
	"time"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) IsValidPhoneNumber(phone string) (bool, string) {
	var message = ""
	var match = true
	match, _ = regexp.MatchString("^[0-9]+$", phone)
	if !match {
		message = "El número de teléfono no es válido"
	}
	if !(len(phone) == 9) {
		match = false
		message = "Longitud de teléfono incorrecta"
	}
	return match, message
}

func (v *UserValidator) IsValidEmail(email string) (bool, string) {
	var message = ""
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`, email)
	if !match {
		message = "El email no es válido"
	}
	return match, message
}

func (v *UserValidator) IsAdult(birthdate string) (bool, string) {
	var message = ""
	var isAdult = true
	parsedBirthdate, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return false, message
	}

	age := time.Now().Year() - parsedBirthdate.Year()
	if !(age >= 18) {
		isAdult = false
		message = "No es mayor de edad"
	}

	return isAdult, message
}

func (v *UserValidator) IsValidDNI(dni string) (bool, string) {
	var message = ""
	var match = true
	match, _ = regexp.MatchString("^[0-9]+$", dni)
	if !match {
		message = "El número de DNO no es válido"
	}
	if !(len(dni) == 8) {
		match = false
		message = "Longitud de DNI incorrecta"
	}
	return match, message
}
