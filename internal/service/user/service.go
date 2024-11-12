package user

import (
	"context"
	"fmt"

	"github.com/sparkymat/nexus/internal/dbx"
)

func New(db *dbx.Queries) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db *dbx.Queries
}

func (s *Service) CreateUser(ctx context.Context, name string, email string, encryptedPassword string) (dbx.User, error) {
	user, err := s.db.CreateUser(ctx, dbx.CreateUserParams{
		Name:              name,
		Email:             email,
		EncryptedPassword: encryptedPassword,
	})
	if err != nil {
		return dbx.User{}, fmt.Errorf("unable to fetch user. err: %w", err)
	}

	return user, nil
}

func (s *Service) FetchUserByEmail(ctx context.Context, email string) (dbx.User, error) {
	user, err := s.db.FetchUserByEmail(ctx, email)
	if err != nil {
		return user, fmt.Errorf("unable to fetch user. err: %w", err)
	}

	return user, nil
}
