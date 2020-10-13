package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Database *Database `envconfig:"DB"`
	Server   *Server   `envconfig:"SERVER"`
}

func NewAppConfig(files ...string) (*App, error) {
	_ = godotenv.Load(files...)

	var conf App
	if err := envconfig.Process("APP", &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
