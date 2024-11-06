package photo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/sparkymat/nexus/internal/tasks"
)

func (s *Service) RescanFolders(ctx context.Context) error {
	for _, folder := range s.photoFolders {
		err := s.scanFolder(ctx, folder)
		if err != nil {
			log.Error(err)

			continue
		}
	}

	return nil
}

func (s *Service) scanFolder(ctx context.Context, path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read dir '%s': %w", path, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			err = s.scanFolder(ctx, filepath.Join(path, entry.Name()))
			if err != nil {
				return fmt.Errorf("failed to read subdir '%s': %w", filepath.Join(path, entry.Name()), err)
			}

			continue
		}

		readFileTask, err := tasks.NewReadPhotoFileTask(filepath.Join(path, entry.Name()))
		if err != nil {
			log.Error(err)

			continue
		}

		taskInfo, err := s.asynqClient.Enqueue(readFileTask)
		if err != nil {
			log.Error(err)
		} else {
			log.Infof("queued task %s for reading file '%s': queue=%s", taskInfo.ID, filepath.Join(path, entry.Name()), taskInfo.Queue)
		}
	}

	return nil
}
