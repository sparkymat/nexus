package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/auth"
	"github.com/sparkymat/nexus/internal/handler"
)

func registerWebRoutes(app *echo.Group, cfg internal.ConfigService, s internal.Services) {
	webApp := app.Group("")

	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", handler.Login(cfg, s.User))
	webApp.POST("/login", handler.DoLogin(cfg, s.User))

	if !cfg.DisableRegistration() {
		webApp.GET("/register", handler.Register(cfg, s.User))
		webApp.POST("/register", handler.DoRegister(cfg, s.User))
	}

	authenticatedWebApp := webApp.Group("")

	if cfg.ReverseProxyAuthentication() {
		authenticatedWebApp.Use(auth.ProxyAuthMiddleware(cfg, s.User))
	} else {
		authenticatedWebApp.Use(auth.Middleware(cfg, s.User))
	}

	authenticatedWebApp.GET("/", handler.Home(cfg, s.User))
}
