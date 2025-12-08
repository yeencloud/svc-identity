package ports

import (
	"context"

	contract "github.com/yeencloud/svc-identity/contract/proto"
)

type Usecases interface {
	Register(_ context.Context, in *contract.RegisterObject) (*contract.RegisterResponse, error)
}
