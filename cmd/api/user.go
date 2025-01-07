package api

import (
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/davi-sant/asafe-vault-go/cmd/pkg/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func UserRoutes(r *gin.Engine, userService *services.UserService) {

	admin := os.Getenv("ADMIN")
	password := os.Getenv("PASSWORD")
	userController := controllers.UserController{UserService: userService}
	ur := r.Group("v1/users", gin.BasicAuth(gin.Accounts{admin: password}))
	{
		ur.POST("/", userController.CreateUser)
		ur.GET("/", userController.GetUserBayEmail)
	}
}
