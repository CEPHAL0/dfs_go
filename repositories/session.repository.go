package repositories

import (
	authModels "backend/models/auth"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var SessionDuration = time.Hour * 24 * 5

type ISessionRepository interface {
	Create(userID uuid.UUID) (*authModels.Session, error)
	Delete(session *authModels.Session) error
	SetSession(session *authModels.Session, c *fiber.Ctx) error
}

type SessionRepository struct {
	*gorm.DB
}

func NewSessionRepository(DB *gorm.DB) ISessionRepository {
	return &SessionRepository{DB: DB}
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

func (sessionRepository *SessionRepository) SetSession(sessionToSet *authModels.Session, c *fiber.Ctx) error {
	var Config = session.Config{
		Expiration:     SessionDuration, // Session expiration duration
		CookieHTTPOnly: true,            // Prevents JavaScript access to cookies
		CookieSecure:   true,            // Ensures cookies are sent only over HTTPS
		CookieSameSite: "Strict",        // Restricts cookies to same-site requests
		CookieName:     "sessionID",     // Custom cookie name
		CookiePath:     "/",             // Cookie path scope
	}

	var store = session.New(Config)

	sess, _ := store.Get(c)

	sess.Set("sessionID", &sessionToSet.SessionID)
	err := sess.Save()

	if err != nil {
		return err
	}

	return nil
}

func (sessionRepository *SessionRepository) Delete(session *authModels.Session) error {
	return nil
}

func getExpiry() time.Time {
	return time.Now().Add(SessionDuration)
}
