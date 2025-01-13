package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

type PasswordServiceController struct {
	PasswordServiceRepository *services.PasswordServiceRepository
}

func (c *PasswordServiceController) CreatePassword(ctx *gin.Context) {
	var req services.PasswordCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if req.UserId <= 0 {
		ctx.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "UserId must be greater than zero",
			Data:    nil,
		})
		return
	}

	if len(strings.TrimSpace(req.ServiceName)) <= 0 {
		ctx.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "Service name cannot be empty",
			Data:    nil,
		})
		return
	}

	if len(req.ServicePassword) < 8 {
		ctx.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "Service password must be less than 8",
			Data:    nil,
		})
		return
	}

	if err := c.PasswordServiceRepository.Create(req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.PayLoadResponse{
		Status:  "created",
		Message: "user created successfully!",
		Data:    req,
	})

}

func (s *PasswordServiceController) GetAllPasswords(context *gin.Context) {

	id := context.Query("user_id")

	if strings.TrimSpace(id) == "" {
		context.JSON(http.StatusInternalServerError, models.PayLoadResponse{
			Status:  "error",
			Message: "error User id  is not empty",
			Data:    nil,
		})
		return
	}
	value, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "error 'ID_USER' is not Int64",
			Data:    nil,
		})
		return
	}

	passwords, err := s.PasswordServiceRepository.GetAllPasswords(value)

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.PayLoadResponse{
			Status:  "error",
			Message: "error fetching passwords",
			Data:    nil,
		})
		return
	}

	context.JSON(http.StatusOK, models.PayLoadResponse{
		Status:  "success",
		Message: "Search successful",
		Data:    passwords,
	})

}
func (s *PasswordServiceController) GetPasswordsByServiceName(context *gin.Context) {

	id := context.Query("user_id")
	service_name := context.Query("service_name")
	fmt.Println(id, service_name)
	value, err := strconv.ParseInt(id, 0, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.PayLoadResponse{
			Status:  "error",
			Message: "error 'ID_USER' is not Int64",
			Data:    nil,
		})
		return
	}

	rows, err := s.PasswordServiceRepository.GetPasswordsByServiceName(value, service_name)

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.PayLoadResponse{
			Status:  "error",
			Message: "error fetching passwords",
			Data:    nil,
		})
		return
	}

	context.JSON(http.StatusOK, models.PayLoadResponse{
		Status:  "success",
		Message: "Search password by service name successful",
		Data:    rows,
	})

}
