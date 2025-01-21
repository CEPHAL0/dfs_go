package repositories

import (
	authModels "backend/models/auth"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	Create() (*authModels.Session, error)
	Delete(session *authModels.Session) error
}

type SessionRepository struct {
	*gorm.DB
}

func NewSessionRepository(DB *gorm.DB) ISessionRepository {
	return &SessionRepository{DB: DB}
}

func (sessionRepository *SessionRepository) Create() (*authModels.Session, error) {
	return &authModels.Session{}, nil
}

func (sessionRepository *SessionRepository) Delete(session *authModels.Session) error {
	return nil
}
