package impl

import (
	authModels "backend/models/auth"
	"backend/repositories"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

func (userRepository *userRepositoryImpl) Insert(username string, password string, email string) authModels.User {
	var user authModels.User

	user.Name = username
	user.Password = password
	user.Email = email

	return user
}
