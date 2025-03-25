// internal/repository/user_repository.go
package repository

import "personjs/myapi/internal/models"

type UserRepository struct{}

func (ur *UserRepository) FindUserByID(id int) *models.User {
	return &models.User{ID: id, Name: "John Doe", Email: "john@example.com"}
}
