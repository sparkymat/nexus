package main

import (
	"github.com/hibiken/asynq"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/nexus/internal/config"
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

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{Addr: cfg.RedisURL()}, nil)

	rescanTask, err := tasks.NewRescanFoldersTask()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	entryID, err := scheduler.Register("0 0 * * *", rescanTask)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
