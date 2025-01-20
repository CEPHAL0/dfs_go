package repositories

import (
	"backend/enums"
	authModels "backend/models/auth"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	// GetByID(username string) authModels.User
	GetByEmail(email string) (*authModels.User, error)
	Insert(username string, password string, email string, role enums.Role) (authModels.User, error)
}

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) IUserRepository {
	return &UserRepository{DB: DB}
}

func (userRepository *UserRepository) GetByEmail(email string) (*authModels.User, error) {
	var user authModels.User

	err := userRepository.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) Insert(username string, email string, password string, role enums.Role) (authModels.User, error) {

	existingUser, err := userRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return authModels.User{}, fmt.Errorf("database error: %w", err)
	}

	if existingUser != nil {
		return authModels.User{}, errors.New("user with email already exists")
	}

	user := authModels.User{
		Name:              username,
		Email:             email,
		Password:          password,
		Role:              role,
		IsPasswordExpired: true,
	}

	result := userRepository.DB.Create(&user)

	if result.Error != nil {
		return authModels.User{}, result.Error
	}

	return user, nil

}
