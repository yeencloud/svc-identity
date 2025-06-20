package ports

import (
	"github.com/yeencloud/lib-base/events"
	"github.com/yeencloud/svc-identity/internal/ports/repository"
)

type Ports struct {
	ViewOriginRepo database.ViewOriginRepository
	EventPublisher events.Publisher
}
