package auth

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
	"golang.org/x/crypto/bcrypt"
)

const defaultBcryptCost = 10

func ProxyAuthMiddleware(cfg internal.ConfigService, userService internal.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name := c.Request().Header.Get(cfg.ProxyAuthNameHeader())

			email := c.Request().Header.Get(cfg.ProxyAuthEmailHeader())
			if email == "" {
				return c.String(http.StatusUnauthorized, "email header missing")
			}

			user, err := userService.FetchUserByEmail(c.Request().Context(), email)
			if err == nil {
				c.Set(UserKey, user)

				return next(c)
			}

			password := strings.ReplaceAll(uuid.New().String(), "-", "")

			encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to generate password")
			}

			user, err = userService.CreateUser(c.Request().Context(), name, email, string(encryptedPasswordBytes))
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to add new user")
			}

			c.Set(UserKey, user)

			return next(c)
		}
	}
}
