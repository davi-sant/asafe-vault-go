package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/davi-sant/asafe-vault-go/cmd/api"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/repositories"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"

	"github.com/davi-sant/asafe-vault-go/cmd/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		result := fmt.Sprintf("Erro ao carregar o arquivo .env: %s ", err)
		fmt.Println(result)
	}

	db, err := config.DBConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}(db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	userRepo := repositories.NewPostgresRepository(db)
	userService := services.NewUserService(userRepo)

	passwordRepo := repositories.NewPostgresPasswordRepository(db)
	passwordService := services.NewPasswordServiceRepository(passwordRepo)

	if err := repositories.InitializeDatabase(db); err != nil {
		log.Fatal(err)
	}
	if err := repositories.InitializePasswordDB(db); err != nil {
		log.Fatal(err)
	}

	api.UserRoutes(r, userService)
	api.PasswordRoutes(r, passwordService)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
