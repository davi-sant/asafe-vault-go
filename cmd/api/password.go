package api

import (
	"os"

	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/davi-sant/asafe-vault-go/cmd/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func PasswordRoutes(r *gin.Engine, passwordService *services.PasswordServiceRepository) {

	admin := os.Getenv("ADMIN")
	password := os.Getenv("PASSWORD")
	passwordController := controllers.PasswordServiceController{PasswordServiceRepository: passwordService}
	ur := r.Group("v1/asafe-vault/password", gin.BasicAuth(gin.Accounts{admin: password}))
	{
		ur.POST("", passwordController.CreatePassword)
		ur.GET("", passwordController.GetAllPasswords)
		ur.GET("/search", passwordController.GetPasswordsByServiceName)
	}
}
