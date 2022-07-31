package repo

import (
	"context"

	"example.com/ex_bun/pkg/models"
)

type UserRepo interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
}
