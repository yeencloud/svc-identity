package database

import (
	"context"

	"github.com/yeencloud/svc-identity/internal/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.User, password string) error

	GetAuthByUsername(ctx context.Context, username string) (*domain.AuthInformation, error)
	GetAuthByEmail(ctx context.Context, email string) (*domain.AuthInformation, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}
