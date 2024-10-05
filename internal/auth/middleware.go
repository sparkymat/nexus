package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal/model"
)

const (
	UserKey = "user"
)

const (
	sessionName = "nexus_session"
	tokenKey    = "auth_token"
)

var ErrTokenMissing = errors.New("token missing")

const ClientNameKey = "client_name"

func Middleware(cfg Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(c, cfg)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
}

func APIMiddleware(cfg Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(c, cfg)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
}

func sessionAuth(c echo.Context, cfg Config) error {
	username, err := LoadUsernameFromSession(cfg, c)
	if err != nil {
		return err
	}

	user, err := model.FetchUserByEmail(c.Request().Context(), username)
	if err != nil {
		return err //nolint:wrapcheck
	}

	c.Set(UserKey, user)

	return nil
}
