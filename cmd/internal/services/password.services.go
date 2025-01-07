package services

import (
	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/repositories"
)

type PasswordCreateRequest struct {
	UserId          int64  `json:"user_id" binding:"required"`
	ServiceName     string `json:"service_name" binding:"required"`
	ServiceUserName string `json:"service_user_name" binding:"required"`
	ServicePassword string `json:"service_password" binding:"required"`
}

type PasswordServiceRepository struct {
	repository repositories.PasswordRepository
}

func NewPasswordServiceRepository(repository repositories.PasswordRepository) *PasswordServiceRepository {
	return &PasswordServiceRepository{repository: repository}
}

func (service *PasswordServiceRepository) Create(payload PasswordCreateRequest) error {
	
	password := models.Password{
		UserId:          payload.UserId,
		ServiceName:     payload.ServiceName,
		ServiceUserName: payload.ServiceUserName,
		ServicePassword: payload.ServicePassword,
	}
	return service.repository.Create(password)
}
