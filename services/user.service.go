package services

import (
	database "backend/config"
	"backend/enums"
	authModels "backend/models/auth"
	"backend/repositories"

	"gorm.io/gorm"
)

type IUserService interface {
	RegisterUser(username, email, password string, role enums.Role) (*authModels.User, *authModels.Session, error)
}

type UserService struct {
	UserRepository    repositories.IUserRepository
	SessionRepository repositories.ISessionRepository
}

func NewUserService(userRepository repositories.IUserRepository, sessionRepository repositories.ISessionRepository) IUserService {
	return &UserService{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}

func (userService *UserService) RegisterUser(username, email, password string, role enums.Role) (*authModels.User, *authModels.Session, error) {

	var user *authModels.User
	var session *authModels.Session

	err := database.WithTransaction(func(tx *gorm.DB) error {
		var err error
		user, err = userService.UserRepository.Create(tx, username, email, password, role)
		if err != nil {
			return err
		}

		// Create the session within the same transaction
		session, err = userService.SessionRepository.Create(tx, user.ID)
		if err != nil {
			return err
		}

		

		// Return nil to commit the transaction
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
