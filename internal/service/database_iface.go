package service

import (
	"context"

	"github.com/sparkymat/nexus/internal/dbx"
)

//nolint:interfacebloat
type DatabaseProvider interface {
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
}
