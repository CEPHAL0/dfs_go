package utils

import (
	authModels "backend/models/auth"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var SessionDuration = time.Hour * 24 * 5

func SetSession(tx *gorm.DB, sessionToSet *authModels.Session, c *fiber.Ctx) error {
	var Config = session.Config{
		Expiration:     SessionDuration, // Session expiration duration
		CookieHTTPOnly: true,            // Prevents JavaScript access to cookies
		CookieSecure:   true,            // Ensures cookies are sent only over HTTPS
		CookieSameSite: "Strict",        // Restricts cookies to same-site requests
		CookiePath:     "/",             // Cookie path scope
		KeyLookup:      "cookie:sessionID",
	}

	var store = session.New(Config)
	sess, err := store.Get(c)

	if err != nil {
		return err
	}

	sess.Set("sessionID", &sessionToSet.SessionID)
	err = sess.Save()

	if err != nil {
		return err
	}

	return errors.New("Test error")
}
