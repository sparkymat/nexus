package internal

import (
	"context"

	"github.com/sparkymat/nexus/internal/dbx"
)

type UserService interface {
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	CreateUser(ctx context.Context, name string, email string, encryptedPassword string) (dbx.User, error)
}
