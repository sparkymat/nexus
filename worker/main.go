package main

import (
	"github.com/hibiken/asynq"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/nexus/internal/config"
	"github.com/sparkymat/nexus/internal/database"
	"github.com/sparkymat/nexus/internal/dbx"
	"github.com/sparkymat/nexus/internal/service/photo"
	"github.com/sparkymat/nexus/internal/tasks"
)

const (
	QueueCriticalName     = "critical"
	QueueCriticalPriority = 6
	QueueDefaultName      = "default"
	QueueDefaultPriority  = 3
	QueueLowName          = "low"
	QueueLowPriority      = 1
)

//nolint:funlen
func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(cfg.DatabaseURL())
	if err != nil {
		log.Error(err)
		panic(err)
	}

	db := dbx.New(dbDriver.DB())
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisURL()})

	photoService := photo.New(cfg.PhotoFolders(), db, asynqClient)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.RedisURL()},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10, //nolint:mnd
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				QueueCriticalName: QueueCriticalPriority,
				QueueDefaultName:  QueueDefaultPriority,
				QueueLowName:      QueueLowPriority,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	mux.Handle(tasks.TypeRescanFolders, tasks.NewFoldersRescanner(photoService))
	mux.Handle(tasks.TypeReadPhotoFile, tasks.NewPhotoFileReader(photoService))

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
