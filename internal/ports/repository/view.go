package database

import (
	"context"

	"github.com/yeencloud/svc-identity/internal/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.User) error
}
