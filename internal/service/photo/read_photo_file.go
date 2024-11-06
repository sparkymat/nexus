package photo

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/nexus/internal/dbx"
)

func (s *Service) ReadPhotoFile(ctx context.Context, path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Errorf("file '%s' does not exist: %w", path, err)
			return nil
		}

		log.Errorf("failed to read file '%s': %w", path, err)
		return nil
	}

	photo, err := s.db.FetchPhotoByPath(ctx, path)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Errorf("failed to check db for '%s': %w", path, err)
		return nil
	}

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		photo, err = s.db.CreatePhoto(ctx, dbx.CreatePhotoParams{
			Path:     path,
			FileType: strings.ToLower(filepath.Ext(path)[1:]),
		})
		if err != nil {
			log.Errorf("failed to create photo entry for '%s': %w", path, err)
			return nil
		}
	}

	log.Infof("add entry '%s' for '%s'", photo.ID.String(), path)

	return nil
}
