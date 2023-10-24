package repository

import (
	"crud-compartamos/models"
	"crud-compartamos/utils"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (s *UserRepository) CreateUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO users (dni, name, last_name, second_last_name, birthdate, phone, email) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.DNI, user.Name, user.LastName, user.SecondLastName, user.Birthdate, user.Phone, user.Email)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepository) IsUserUnique(dni string) bool {
	query := "SELECT COUNT(*) FROM users WHERE dni = ?"
	var count int
	err := s.db.QueryRow(query, dni).Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}

func (s *UserRepository) UserExists(dni string) bool {
	query := "SELECT COUNT(*) FROM users WHERE dni = ?"
	var count int
	err := s.db.QueryRow(query, dni).Scan(&count)
	if err != nil {
		return false
	}
	return count >= 1
}

func (s *UserRepository) UpdateUser(user *models.User, dniParam string) error {
	_, err := s.db.Exec("UPDATE users SET dni = ?, name = ?, last_name = ?, second_last_name = ?, birthdate = ?, phone = ?, email = ? WHERE dni = ?",
		user.DNI, user.Name, user.LastName, user.SecondLastName, user.Birthdate, user.Phone, user.Email, dniParam)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepository) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, dni, name, last_name, second_last_name, birthdate, phone, email FROM users"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.DNI, &user.Name, &user.LastName, &user.SecondLastName, &user.Birthdate, &user.Phone, &user.Email); err != nil {
			return nil, err
		}
		user.Birthdate = utils.NewFormatter().FormatBirthdate(user.Birthdate)

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserRepository) GetUserByDNI(dni string) (*models.User, error) {
	query := "SELECT id, dni, name, last_name, second_last_name, birthdate, phone, email FROM users WHERE dni = ?"
	var user models.User

	if err := s.db.QueryRow(query, dni).Scan(&user.ID, &user.DNI, &user.Name, &user.LastName, &user.SecondLastName, &user.Birthdate, &user.Phone, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) DeleteUserByDNI(dni string) (int64, error) {
	result, err := s.db.Exec("DELETE FROM users WHERE dni = ?", dni)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
