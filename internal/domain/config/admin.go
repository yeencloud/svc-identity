package config

import "github.com/yeencloud/lib-shared/config"

type Admin struct {
	AdministratorUser     config.Secret `config:"ADMIN_USER" default:"admin"`
	AdministratorPassword config.Secret `config:"ADMIN_PASSWORD" default:"admin"`
}
