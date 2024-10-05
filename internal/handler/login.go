package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal/model"
	"github.com/sparkymat/nexus/internal/view"
)

func Login(cfg Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderLoginPage(cfg, c, "", "")
	}
}

func DoLogin(cfg Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		err := model.UserLogin(c, email, password)
		if err != nil {
			log.Printf("failed to log user in: %s", err.Error())

			return renderLoginPage(cfg, c, email, "Authentication failed")
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func renderLoginPage(cfg Config, c echo.Context, email string, errorMessage string) error {
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
