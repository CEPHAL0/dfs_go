package repositories

import (
	database "backend/config"
	authModels "backend/models/auth"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var SessionDuration = time.Hour * 24 * 5

type ISessionRepository interface {
	Create(tx *gorm.DB, userID uuid.UUID) (*authModels.Session, error)
	Delete(session *authModels.Session) error
}

type SessionRepository struct {
	*gorm.DB
}

func NewSessionRepository() ISessionRepository {
	return &SessionRepository{DB: database.GetDB()}
}

func (sessionRepository *SessionRepository) Create(tx *gorm.DB, userID uuid.UUID) (*authModels.Session, error) {
	session := authModels.Session{
		SessionID: uuid.New().String(),
		UserID:    userID,
		Expires:   getExpiry(),
	}

	err := tx.Create(&session).Error

	if err != nil {
		return &authModels.Session{}, err
	}

	return &session, nil

}

func (sessionRepository *SessionRepository) Delete(session *authModels.Session) error {
	return nil
}

func getExpiry() time.Time {
	return time.Now().Add(SessionDuration)
}
