package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeReadPhotoFile = "photo:read_file"
)

type ReadPhotoFilePayload struct {
	Path string `json:"path"`
}

func NewReadPhotoFileTask(path string) (*asynq.Task, error) {
	payload, err := json.Marshal(ReadPhotoFilePayload{
		Path: path,
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return asynq.NewTask(TypeReadPhotoFile, payload), nil
}

type PhotoFileReader struct {
	photoService photoService
}

func (f *PhotoFileReader) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ReadPhotoFilePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w: %w", err, asynq.SkipRetry)
	}

	// Do stuff
	return f.photoService.ReadPhotoFile(ctx, p.Path) //nolint:wrapcheck
}

func NewPhotoFileReader(photoService photoService) *PhotoFileReader {
	return &PhotoFileReader{
		photoService: photoService,
	}
}
