package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type User struct {
}

func FetchUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func CreateUser(ctx context.Context, name, username, password string) (*User, error) {
	return nil, nil
}

func UserLogin(c echo.Context, email, password string) error {
	return nil
}
