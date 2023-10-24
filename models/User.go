package models

type User struct {
	ID             int
	DNI            string `json:"dni"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	SecondLastName string `json:"second_last_name"`
	Birthdate      string `json:"birthdate"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}
