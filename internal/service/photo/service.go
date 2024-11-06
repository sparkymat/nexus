package photo

import (
	"github.com/sparkymat/nexus/internal/service"
)

func New(photoFolders []string, db service.DatabaseProvider, asynqClient service.AsynqProvider) *Service {
	return &Service{
		photoFolders: photoFolders,
		db:           db,
		asynqClient:  asynqClient,
	}
}

type Service struct {
	photoFolders []string
	db           service.DatabaseProvider
	asynqClient  service.AsynqProvider
}
