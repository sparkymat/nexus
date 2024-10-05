package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal/config"
	"github.com/sparkymat/nexus/internal/route"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	route.Setup(e, cfg)
	route.PrintRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
