package service

import (
	"context"

	contract "github.com/yeencloud/svc-identity/contract/proto"
)

func (s service) Register(_ context.Context, in *contract.RegisterObject) (*contract.RegisterResponse, error) {
	return &contract.RegisterResponse{Id: in.Mail}, nil
}
