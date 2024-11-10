package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type photoService interface {
	RescanFolders(ctx context.Context) error
	ReadPhotoFile(ctx context.Context, path string) error
}

const (
	TypeRescanFolders = "photo:rescan_folders"
)

type RescanFoldersPayload struct{}

func NewRescanFoldersTask() (*asynq.Task, error) {
	payload, err := json.Marshal(RescanFoldersPayload{})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return asynq.NewTask(TypeRescanFolders, payload), nil
}

type FoldersRescanner struct {
	photoService photoService
}

func (f *FoldersRescanner) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p RescanFoldersPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w: %w", err, asynq.SkipRetry)
	}

	// Do stuff
	return f.photoService.RescanFolders(ctx) //nolint:wrapcheck
}

func NewFoldersRescanner(photoService photoService) *FoldersRescanner {
	return &FoldersRescanner{
		photoService: photoService,
	}
}
