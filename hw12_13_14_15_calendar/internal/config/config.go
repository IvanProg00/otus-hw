package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DatabaseConf `yaml:"db"`
	Logger   LoggerConf   `yaml:"logger"`
	Server   ServerConf   `yaml:"server"`
}

type DatabaseConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfigFromYaml(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
