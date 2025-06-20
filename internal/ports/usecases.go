package ports

import (
	"context"

	"github.com/yeencloud/svc-identity/internal/domain"
)

type Usecases interface {
	Viewed(ctx context.Context, origin domain.ViewOrigin) ([]domain.ViewOrigin, error)
}
