package controllers

import (
	"crud-compartamos/models"
	"crud-compartamos/repository"
	"crud-compartamos/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context, userRepository *repository.UserRepository) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isUnique := userRepository.IsUserUnique(user.DNI)
	if !isUnique {
		errorMessage := map[string]string{
			"error": "DNI ya se encuentra registrado",
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	validate, message := isUserVal(user)

	if !validate {
		errorMessage := map[string]string{
			"error": message,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return

	}

	if err := userRepository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func GetUsers(c *gin.Context, userRepository *repository.UserRepository) {
	users, err := userRepository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	formattedResult := utils.NewFormatter().FormatUsersResult(users)
	c.JSON(http.StatusOK, formattedResult)
}

func GetUserByDni(c *gin.Context, userRepository *repository.UserRepository) {
	dni := c.Param("id")

	user, err := userRepository.GetUserByDNI(dni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuario por DNI"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	user.Birthdate = utils.NewFormatter().FormatBirthdate(user.Birthdate)

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context, userRepository *repository.UserRepository) {
	var user models.User
	userDni := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists := userRepository.UserExists(userDni)
	if !exists {
		errorMessage := map[string]string{
			"error": "El usuario no existe",
		}
		c.JSON(http.StatusNotFound, errorMessage)
		return
	}
	validate, message := isUserVal(user)

	if !validate {
		errorMessage := map[string]string{
			"error": message,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return

	}
	if err := userRepository.UpdateUser(&user, userDni); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func DeleteUser(c *gin.Context, userRepository *repository.UserRepository) {
	userDni := c.Param("id")
	rowAffected, err := userRepository.DeleteUserByDNI(userDni)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if rowAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No existe usuario con ese DNI"})
		return
	}

}

func isUserVal(user models.User) (bool, string) {
	userValidator := utils.NewUserValidator()

	var validate = true
	var message = ""

	validate, message = userValidator.IsValidPhoneNumber(user.Phone)
	if !validate {
		return false, message
	}

	validate, message = userValidator.IsValidEmail(user.Email)
	if !validate {
		return false, message
	}

	validate, message = userValidator.IsAdult(user.Birthdate)
	if !validate {
		return false, message
	}

	validate, message = userValidator.IsValidDNI(user.DNI)
	if !validate {
		return false, message
	}

	return true, message
}
