package rpc

import (
	"github.com/yeencloud/lib-shared/validation"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
	"github.com/yeencloud/svc-identity/internal/ports"
)

type Handler struct {
	usecases  ports.Usecases
	validator *validation.Validator

	contract.UnimplementedIdentityServiceServer
}

func NewRPCHandler(usecases ports.Usecases, validator *validation.Validator) *Handler {
	return &Handler{
		usecases:  usecases,
		validator: validator,
	}
}
