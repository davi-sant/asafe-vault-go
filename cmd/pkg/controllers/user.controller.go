package controllers

import (
	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func (userController *UserController) CreateUser(context *gin.Context) {

	var req services.UserCreatRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "error when processing user data: " + err.Error(),
			Data:    nil,
		})
		return
	}

	if err := userController.UserService.CreateUser(req); err != nil {
		context.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "error creating user: " + err.Error(),
			Data:    nil,
		})
		return
	}

	context.JSON(http.StatusCreated, models.PayLoadResponse{
		Status:  "created",
		Message: "user created successfully!",
		Data:    req,
	})
}

func (userController *UserController) GetUserBayEmail(context *gin.Context) {
	email := context.Query("email")
	user, err := userController.UserService.GetUserBayEmail(email)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "error getting user: " + err.Error(),
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, models.PayLoadResponse{
		Status:  "success",
		Message: "user found successfully",
		Data:    user,
	})
}
