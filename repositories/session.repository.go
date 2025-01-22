package repositories

import (
	authModels "backend/models/auth"
	"time"

	database "backend/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var SessionDuration = time.Hour * 24 * 5

type ISessionRepository interface {
	Create(userID uuid.UUID) (*authModels.Session, error)
	CreateWithTx(userID uuid.UUID, tx *gorm.DB) (*authModels.Session, error)
	Delete(session *authModels.Session) error
}

type SessionRepository struct {
	*gorm.DB
}

func NewSessionRepository() ISessionRepository {
	return &SessionRepository{DB: database.GetDB()}
}

func (sessionRepository *SessionRepository) Create(userID uuid.UUID) (*authModels.Session, error) {
	session := authModels.Session{
		SessionID: uuid.New().String(),
		UserID:    userID,
		Expires:   getExpiry(),
	}

	sess := sessionRepository.DB.Create(&session)

	if sess.Error != nil {
		return &authModels.Session{}, sess.Error
	}

	return &session, nil

}

func (sessionRepository *SessionRepository) CreateWithTx(userID uuid.UUID, tx *gorm.DB) (*authModels.Session, error) {
	session := authModels.Session{
		SessionID: uuid.New().String(),
		UserID:    userID,
		Expires:   getExpiry(),
	}

	sess := tx.Create(&session)

	if sess.Error != nil {
		return &authModels.Session{}, sess.Error
	}

	return &session, nil

}

func (sessionRepository *SessionRepository) Delete(session *authModels.Session) error {
	return nil
}

func getExpiry() time.Time {
	return time.Now().Add(SessionDuration)
}
