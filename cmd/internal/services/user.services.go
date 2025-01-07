package services

import (
	"errors"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/repositories"
)

type UserCreatRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) CreateUser(req UserCreatRequest) error {
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	existUser, _ := service.repository.GetUserBayEmail(req.Email)
	if existUser != nil {
		return errors.New("user with this email already exists")
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	return service.repository.Create(user)
}

func (service *UserService) GetUserBayEmail(email string) (*models.User, error) {
	user, err := service.repository.GetUserBayEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
