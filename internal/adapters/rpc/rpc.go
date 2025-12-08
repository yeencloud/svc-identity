package rpc

import (
	contract "github.com/yeencloud/svc-identity/contract/proto"
	"github.com/yeencloud/svc-identity/internal/ports"
)

type RPCServerResponder struct {
	usecases ports.Usecases

	contract.UnimplementedIdentityServiceServer
}

func NewRPCServer(usecases ports.Usecases) *RPCServerResponder {
	return &RPCServerResponder{
		usecases: usecases,
	}
}
