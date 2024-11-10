package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/view"
)

func Home(cfg internal.ConfigService, _ internal.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			return c.String(http.StatusInternalServerError, "server error")
		}

		pageHTML := view.Home()
		document := view.BasicLayout(cfg.Version(), "nexus", csrfToken, pageHTML)

		return Render(c, http.StatusOK, document)
	}
}
