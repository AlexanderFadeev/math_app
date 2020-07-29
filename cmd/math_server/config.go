package main

import "github.com/kelseyhightower/envconfig"

type config struct {
	APIPort uint `envconfig:"api_port" required:"true"`
}

func loadConfig() (config, error) {
	var cfg config

	err := envconfig.Process("math_app", &cfg)
	if err != nil {
		return config{}, err
	}

	return cfg, nil
}
