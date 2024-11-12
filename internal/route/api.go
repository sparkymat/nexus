package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/auth"
	"github.com/sparkymat/nexus/internal/handler/api"
)

//nolint:funlen
func registerAPIRoutes(app *echo.Group, cfg internal.ConfigService, services internal.Services) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, services.User))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, services.User))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))

	apiGroup.POST("/objects", api.ObjectsCreate(services))
}
