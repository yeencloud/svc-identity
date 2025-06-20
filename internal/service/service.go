package service

import (
	"github.com/yeencloud/lib-base/events"
	"github.com/yeencloud/svc-identity/internal/ports"
	"github.com/yeencloud/svc-identity/internal/ports/repository"
)

type service struct {
	ports *ports.Ports
}

func NewUsecases(viewRepository database.ViewOriginRepository, eventPublisher events.Publisher) service {
	return service{
		ports: &ports.Ports{
			ViewOriginRepo: viewRepository,

			EventPublisher: eventPublisher,
		},
	}
}
