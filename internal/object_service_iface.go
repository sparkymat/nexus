package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/sparkymat/nexus/internal/dbx"
	"github.com/sparkymat/nexus/internal/service/object"
)

type ObjectService interface {
	CreateObject(ctx context.Context, options object.CreateObjectOptions) (dbx.Object, []dbx.Property, error)
	CreateProperty(ctx context.Context, options object.CreatePropertyOptions) (dbx.Property, *dbx.Object, error)
	FetchObject(ctx context.Context, id uuid.UUID) (dbx.Object, []dbx.Property, map[uuid.UUID]dbx.Object, error)
}
