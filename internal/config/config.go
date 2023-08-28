package config

import (
	"github.com/kelseyhightower/envconfig"

	"gymshark-test/pkg/errors"
)

var MainConfig *Config

type Config struct{}

// InitConfigs Initialize service configuration from ENV
func InitConfigs() error {
	var cfg Config

	if err := envconfig.Process("APP_CONFIG", &cfg); err != nil {
		return errors.ENVReadError.SetDevMessage(err.Error())
	}

	MainConfig = &cfg
	return nil
}
