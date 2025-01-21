package services

import (
	"backend/config"
	"backend/enums"
	authModels "backend/models/auth"
	"backend/repositories"
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

	tx := config.DB.Begin()

	user, err := userService.UserRepository.Insert(username, email, password, role)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	session, err := userService.SessionRepository.Create(user.ID)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	tx.Commit()
	return user, session, nil
}
