package rpc

import (
	"context"
	"errors"

	contract "github.com/yeencloud/svc-identity/contract/proto"
)

func (server *RPCServerResponder) Register(ctx context.Context, object *contract.RegisterObject) (*contract.RegisterResponse, error) {
	return nil, errors.New("NOT IMPLEMENTED")
	// return server.usecases.Register(ctx, object)
}
