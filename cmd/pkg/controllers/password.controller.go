package controllers

import (
	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

func (c *PasswordServiceController) GetPassword(context *gin.Context) {}
