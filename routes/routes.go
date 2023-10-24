package routes

import (
	"crud-compartamos/controllers"
	"crud-compartamos/repository"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, userRepository *repository.UserRepository) {
	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, userRepository)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		controllers.GetUserByDni(c, userRepository)
	})

	r.POST("/user", func(c *gin.Context) {
		controllers.CreateUser(c, userRepository)
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		controllers.UpdateUser(c, userRepository)
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		controllers.DeleteUser(c, userRepository)
	})
}
