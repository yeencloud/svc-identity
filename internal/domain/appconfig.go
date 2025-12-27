package domain

import "github.com/yeencloud/svc-identity/internal/domain/config"

type AppConfig struct {
	Admin          config.Admin
	Authentication config.Authentication
	Registration   config.Registration
}
