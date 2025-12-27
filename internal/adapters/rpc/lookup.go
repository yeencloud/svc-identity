package rpc

import (
	"context"

	"github.com/yeencloud/lib-shared/apperr"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
)

func (server *Handler) LookupUserIDByUsername(context.Context, *contract.IDLookupRequest) (*contract.IDLookupResponse, error) {
	return nil, apperr.NotImplementedError{}
}

func (server *Handler) LookupUserIDByEmail(context.Context, *contract.IDLookupRequest) (*contract.IDLookupResponse, error) {
	return nil, apperr.NotImplementedError{}
}

func (server *Handler) LookupUserProfileByID(context.Context, *contract.ProfileLookupRequest) (*contract.ProfileLookupResponse, error) {
	return nil, apperr.NotImplementedError{}
}
