package rpc

import (
	"context"

	"github.com/yeencloud/lib-shared/apperr"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
)

func (server *Handler) Authenticate(context.Context, *contract.AuthRequest) (*contract.AuthResponse, error) {
	return nil, apperr.NotImplementedError{}
}
