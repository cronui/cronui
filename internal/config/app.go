package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type App struct {
	Server *Server
}

func NewAppConfig(file string) (*App, error) {
	var conf App
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(yamlFile, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
