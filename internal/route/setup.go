package route

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/gorilla/sessions"
	"github.com/jedib0t/go-pretty/table"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"github.com/sparkymat/nexus/internal/auth"
	"github.com/sparkymat/nexus/internal/handler"
)

type Cruddable interface {
	Create(c echo.Context) error
	Destroy(c echo.Context) error
	Index(c echo.Context) error
	Show(c echo.Context) error
	Update(c echo.Context) error
}

type Config interface {
	DisableRegistration() bool
	JWTSecret() string
	SessionSecret() string
	ReverseProxyAuthentication() bool
	ProxyAuthNameHeader() string
	ProxyAuthEmailHeader() string
}

func Setup(e *echo.Echo, cfg Config) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app := e.Group("")
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] Got ${method} on ${uri} from ${remote_ip}. Responded with ${status} in ${latency_human}.\n",
	}))
	app.Use(middleware.Recover())

	app.Static("/js", "public/js")
	app.Static("/css", "public/css")
	app.Static("/images", "public/images")
	app.Static("/fonts", "public/fonts")

	app.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret()))))

	webApp := app.Group("")

	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", handler.Login(cfg))
	webApp.POST("/login", handler.DoLogin(cfg))

	if !cfg.DisableRegistration() {
		webApp.GET("/register", handler.Register())
		webApp.POST("/register", handler.DoRegister())
	}

	authenticatedWebApp := webApp.Group("")

	if cfg.ReverseProxyAuthentication() {
		authenticatedWebApp.Use(auth.ProxyAuthMiddleware(cfg))
	} else {
		authenticatedWebApp.Use(auth.Middleware(cfg))
	}

	authenticatedWebApp.GET("/", handler.Home())
	addResource(webApp, "subject", &handler.Subject{})
}

func addResource(e *echo.Group, name string, cruddable Cruddable) {
	name = strings.ToLower(name)

	plural := pluralize.NewClient().Plural(name)

	e.GET("/"+name+"/:id", cruddable.Show)
	e.GET("/"+plural, cruddable.Index)
	e.POST("/"+plural, cruddable.Create)
	e.PATCH("/"+name+"/:id", cruddable.Update)
	e.DELETE("/"+name+"/:id", cruddable.Update)
}

func PrintRoutes(e *echo.Echo) {
	routes := e.Routes()
	slices.SortFunc(routes, func(a, b *echo.Route) int {
		return strings.Compare(a.Path, b.Path)
	})

	routeRows := lo.Map(routes, func(r *echo.Route, _ int) table.Row {
		method := r.Method
		if method == echo.RouteNotFound {
			method = "ANY"
		}

		return table.Row{
			method + " " + r.Path,
			r.Name,
		}
	})

	tw := table.NewWriter()
	tw.AppendRows(routeRows)
	tw.SetIndexColumn(1)
	tw.SetTitle("Available Routes")
	tw.Style().Options.SeparateRows = true
	fmt.Println(tw.Render())
}
