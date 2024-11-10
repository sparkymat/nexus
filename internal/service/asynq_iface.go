package service

import "github.com/hibiken/asynq"

type AsynqProvider interface {
	Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}
