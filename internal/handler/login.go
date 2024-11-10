package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/auth"
	"github.com/sparkymat/nexus/internal/view"
	"golang.org/x/crypto/bcrypt"
)

func Login(cfg internal.ConfigService, _ internal.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderLoginPage(cfg, c, "", "")
	}
}

func DoLogin(cfg internal.ConfigService, userService internal.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		user, err := userService.FetchUserByEmail(c.Request().Context(), email)
		if err != nil {
			log.Printf("failed to load user. err: %v", err)

			return renderLoginPage(cfg, c, email, "Authentication failed")
		}

		if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) != nil {
			log.Printf("password does not match")

			return renderLoginPage(cfg, c, email, "Authentication failed")
		}

		err = auth.SaveEmailToSession(cfg, c, user.Email)
		if err != nil {
			log.Printf("failed to save email to session. err: %v", err)

			return renderLoginPage(cfg, c, email, "Authentication failed")
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func renderLoginPage(cfg internal.ConfigService, c echo.Context, email string, errorMessage string) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "server error")
	}

	pageHTML := view.Login(cfg.DisableRegistration(), csrfToken, email, errorMessage)
	document := view.Layout("nexus | login", csrfToken, pageHTML)

	//nolint:wrapcheck
	return Render(c, http.StatusOK, document)
}
