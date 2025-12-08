package service

import (
	"github.com/yeencloud/lib-base/events"
	"github.com/yeencloud/svc-identity/internal/ports"
	"github.com/yeencloud/svc-identity/internal/ports/repository"
)

type service struct {
	ports *ports.Ports
}

func NewUsecases(userRepository database.UserRepository, eventPublisher events.Publisher) service {
	return service{
		ports: &ports.Ports{
			EventPublisher: eventPublisher,
		},
	}
}
