package utils

import (
	authModels "backend/models/auth"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionDuration = time.Hour * 24 * 5

func SetSession(sessionToSet *authModels.Session, c *fiber.Ctx) error {
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

	return err
}
