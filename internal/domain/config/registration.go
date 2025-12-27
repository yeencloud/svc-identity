package config

type Registration struct {
	Enabled bool `config:"REGISTRATION_ENABLED" default:"true"`
}
