package config

type Authentication struct {
	PermitAdminLogin bool `config:"PERMIT_ADMIN_LOGIN" default:"false"`
}
