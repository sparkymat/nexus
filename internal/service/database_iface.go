package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sparkymat/nexus/internal/dbx"
)

//nolint:interfacebloat
type DatabaseProvider interface {
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	CreatePhoto(ctx context.Context, arg dbx.CreatePhotoParams) (dbx.Photo, error)
	FetchPhotoByID(ctx context.Context, id uuid.UUID) (dbx.Photo, error)
	FetchPhotoByPath(ctx context.Context, path string) (dbx.Photo, error)
}
