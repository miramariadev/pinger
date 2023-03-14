package config

import (
	"github.com/asaskevich/govalidator"
)

type Config struct {
	Pinger      PingerInfo `toml:"pinger"`
	Environment string     `valid:"required"`
}

type PingerInfo struct {
	Addr string `valid:"required"`
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) IsValid() error {
	_, err := govalidator.ValidateStruct(config)
	if err != nil {
		return err
	}
	return nil
}
