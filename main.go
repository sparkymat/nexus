package main

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/config"
	"github.com/sparkymat/nexus/internal/database"
	"github.com/sparkymat/nexus/internal/dbx"
	"github.com/sparkymat/nexus/internal/route"
	"github.com/sparkymat/nexus/internal/service/user"
	"github.com/sparkymat/nexus/internal/tasks"
)

var Version = "development"

func main() {
	cfg, err := config.New(Version)
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(cfg.DatabaseURL())
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		log.Error(err)
		panic(err)
	}

	// Initialize web server
	db := dbx.New(dbDriver.DB())

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisURL()})
	defer asynqClient.Close()

	rescanTask, err := tasks.NewRescanFoldersTask()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	taskInfo, err := asynqClient.Enqueue(rescanTask, asynq.Unique(time.Hour))
	if err != nil {
		log.Error(err) // Don't panic
	} else {
		log.Infof("queued task %s for rescanning folders: queue=%s", taskInfo.ID, taskInfo.Queue)
	}

	userService := user.New(db)

	services := internal.Services{
		User: userService,
	}

	e := echo.New()

	route.Setup(e, cfg, services)
	route.PrintRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
