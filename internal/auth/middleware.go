package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
)

const (
	UserKey = "user"
)

const (
	sessionName = "scenes_session"
	tokenKey    = "auth_token"
)

var ErrTokenMissing = errors.New("token missing")

const ClientNameKey = "client_name"

func Middleware(cfg internal.ConfigService, userService internal.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(c, cfg, userService)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
}

func APIMiddleware(cfg internal.ConfigService, userService internal.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(c, cfg, userService)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
}

func sessionAuth(c echo.Context, cfg internal.ConfigService, userService internal.UserService) error {
	email, err := LoadEmailFromSession(cfg, c)
	if err != nil {
		return err
	}

	user, err := userService.FetchUserByEmail(c.Request().Context(), email)
	if err != nil {
		return err //nolint:wrapcheck
	}

	c.Set(UserKey, user)

	return nil
}
