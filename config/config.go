package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Env      string `yaml:"env"`
	Database struct {
		MySQL struct {
			User     string `yaml:"user"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
			Password string `yaml:"password"`
		} `yaml:"mysql"`
	} `yaml:"database"`
	App struct {
		PageSize int64 `yaml:"page_size"`
	} `yaml:"app"`
}

func LoadConfig() (*Config, error) {
	path := os.Getenv("APP_CONFIG_PATH")
	if path == "" {
		return nil, errors.New("env [APP_CONFIG_PATH] not defined")
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err := yaml.Unmarshal(file, c); err != nil {
		return nil, err
	}
	return c, nil
}
