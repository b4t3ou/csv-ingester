package config

import (
	"github.com/caarlos0/env"
	"gopkg.in/go-playground/validator.v9"
)

// Config is the config struct
type Config struct {
	Port        string `env:"PORT" validate:"required"`
	ServerHost  string `env:"SERVER_HOST" validate:"required"`
	ServiceType string `env:"SERVICE_TYPE" validate:"required"`
}

// New returns with a new Cfg object
func New() (*Config, error) {
	c := &Config{}

	err := env.Parse(c)
	if err != nil {
		return c, err
	}

	v := validator.New()

	err = v.Struct(c)

	return c, err
}
