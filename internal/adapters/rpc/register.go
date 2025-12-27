package rpc

import (
	"context"

	"github.com/mailgun/errors"
	"github.com/yeencloud/lib-shared/apperr"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
	"github.com/yeencloud/svc-identity/internal/domain"
)

func (server *Handler) Register(ctx context.Context, object *contract.RegisterObject) (*contract.RegisterResponse, error) {
	if object == nil {
		//TODO: Create an error for that
		return nil, errors.Wrap(apperr.InvalidArgumentError{}, "input is nil")
	}

	params := domain.CreateUserParams{
		Mail:     object.Email,
		Username: object.Username,
		Password: object.Password,
	}

	err := server.validator.ValidateStruct(params)
	if err != nil {
		return nil, err
	}

	return server.usecases.Register(ctx, params)
}
