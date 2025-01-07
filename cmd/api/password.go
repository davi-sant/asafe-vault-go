package api

import (
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/davi-sant/asafe-vault-go/cmd/pkg/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func PasswordRoutes(r *gin.Engine, passwordService *services.PasswordServiceRepository) {

	admin := os.Getenv("ADMIN")
	password := os.Getenv("PASSWORD")
	passwordController := controllers.PasswordServiceController{PasswordServiceRepository: passwordService}
	ur := r.Group("v1/password", gin.BasicAuth(gin.Accounts{admin: password}))
	{
		ur.POST("/", passwordController.CreatePassword)
		ur.GET("/", passwordController.GetPassword)
	}
}
