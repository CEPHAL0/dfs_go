package services

import (
	"backend/enums"
	authModels "backend/models/auth"
	"backend/repositories"

	"gorm.io/gorm"
)

type IUserService interface {
	RegisterUser(username, email, password string, role enums.Role, tx *gorm.DB) (*authModels.User, *authModels.Session, error)
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

func (userService *UserService) RegisterUser(username, email, password string, role enums.Role, tx *gorm.DB) (*authModels.User, *authModels.Session, error) {

	user, err := userService.UserRepository.CreateWithTx(username, email, password, role, tx)
	if err != nil {
		return nil, nil, err
	}

	session, err := userService.SessionRepository.CreateWithTx(user.ID, tx)
	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
