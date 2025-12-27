package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
	"github.com/yeencloud/svc-identity/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s service) Register(ctx context.Context, in domain.CreateUserParams) (*contract.RegisterResponse, error) {
	return handleWithTransaction[contract.RegisterResponse](ctx, s.transaction, func(dbctx context.Context) (*contract.RegisterResponse, error) {
		if !s.appConfig.Registration.Enabled {
			return nil, domain.DisabledRegistrationError{}
		}

		newUser := domain.User{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			Email:     in.Mail,
			Username:  in.Username,
		}

		password := in.Password
		fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		err = s.models.user.AddUser(dbctx, newUser, string(fromPassword))
		if err != nil {
			return nil, err
		}

		return &contract.RegisterResponse{
			Id: newUser.ID,
		}, nil
	})
}
