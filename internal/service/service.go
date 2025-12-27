package service

import (
	"github.com/yeencloud/lib-base/events"
	"github.com/yeencloud/lib-base/transaction"
	"github.com/yeencloud/svc-identity/internal/domain"
	"github.com/yeencloud/svc-identity/internal/ports"
	"github.com/yeencloud/svc-identity/internal/ports/repository"
)

type models struct {
	user database.UserRepository
}

type service struct {
	ports     *ports.Ports
	appConfig domain.AppConfig
	models    models

	transaction transaction.TransactionInterface
}

func NewUsecases(userRepository database.UserRepository, appConfig domain.AppConfig, eventPublisher events.Publisher, transaction transaction.TransactionInterface) service {
	return service{
		ports: &ports.Ports{
			EventPublisher: eventPublisher,
		},
		appConfig: appConfig,
		models: models{
			user: userRepository,
		},
		transaction: transaction,
	}
}
