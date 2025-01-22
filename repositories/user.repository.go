package repositories

import (
	database "backend/config"
	"backend/enums"
	authModels "backend/models/auth"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// GetByID(username string) authModels.User
	GetByEmail(email string) (*authModels.User, error)
	Create(username string, password string, email string, role enums.Role) (*authModels.User, error)
	CreateWithTx(username string, password string, email string, role enums.Role, tx *gorm.DB) (*authModels.User, error)
	ComparePassword(toCompare string, original string) bool
}

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository() IUserRepository {
	return &UserRepository{DB: database.GetDB()}
}

func (userRepository *UserRepository) GetByEmail(email string) (*authModels.User, error) {
	var user authModels.User

	err := userRepository.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) Create(username string, email string, password string, role enums.Role) (*authModels.User, error) {

	existingUser, err := userRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &authModels.User{}, fmt.Errorf("database error: %w", err)
	}

	if existingUser != nil {
		return &authModels.User{}, errors.New("user with email already exists")
	}

	pwdByte := []byte(password)

	hashedPassword, err := hashPassword(pwdByte)

	if err != nil {
		return &authModels.User{}, err
	}

	user := authModels.User{
		Name:              username,
		Email:             email,
		Password:          hashedPassword,
		Role:              role,
		IsPasswordExpired: true,
	}

	result := userRepository.DB.Create(&user)

	if result.Error != nil {
		return &authModels.User{}, result.Error
	}

	return &user, nil

}

func (userRepository *UserRepository) CreateWithTx(username string, email string, password string, role enums.Role, tx *gorm.DB) (*authModels.User, error) {

	existingUser, err := userRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if existingUser != nil {
		return &authModels.User{}, errors.New("user with email already exists")
	}

	pwdByte := []byte(password)

	hashedPassword, err := hashPassword(pwdByte)
	if err != nil {
		return nil, err
	}

	user := authModels.User{
		Name:              username,
		Email:             email,
		Password:          hashedPassword,
		Role:              role,
		IsPasswordExpired: true,
	}

	result := tx.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (userRepository *UserRepository) ComparePassword(toCompare string, hashed string) bool {

	byteHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHashed, []byte(toCompare))

	if err != nil {
		return false
	} else {
		return true
	}
}

func hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
