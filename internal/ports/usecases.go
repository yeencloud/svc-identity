package ports

import (
	"context"

	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
	"github.com/yeencloud/svc-identity/internal/domain"
)

type Usecases interface {
	Register(_ context.Context, in domain.CreateUserParams) (*contract.RegisterResponse, error)
}
