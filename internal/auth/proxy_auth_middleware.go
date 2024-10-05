package auth

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal/model"
	"golang.org/x/crypto/bcrypt"
)

const defaultBcryptCost = 10

func ProxyAuthMiddleware(cfg Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name := c.Request().Header.Get(cfg.ProxyAuthNameHeader())

			email := c.Request().Header.Get(cfg.ProxyAuthEmailHeader())
			if email == "" {
				return c.String(http.StatusUnauthorized, "user header missing")
			}

			user, err := model.FetchUserByEmail(c.Request().Context(), email)
			if err == nil {
				c.Set(UserKey, user)

				return next(c)
			}

			password := strings.ReplaceAll(uuid.New().String(), "-", "")

			encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to generate password")
			}

			user, err = model.CreateUser(c.Request().Context(), name, email, string(encryptedPasswordBytes))
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to add new user")
			}

			c.Set(UserKey, user)

			return next(c)
		}
	}
}
