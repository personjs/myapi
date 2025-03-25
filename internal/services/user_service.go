// internal/services/user_service.go
package services

import (
	"personjs/myapi/internal/models"
	"personjs/myapi/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (us *UserService) GetUser(id int) *models.User {
	return us.Repo.FindUserByID(id)
}
