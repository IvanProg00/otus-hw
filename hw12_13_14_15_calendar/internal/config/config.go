package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Database DatabaseConf `yaml:"db"`
	Logger   LoggerConf   `yaml:"logger"`
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

func NewConfigFromYaml(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
