package configuration

import (
	"errors"
	"os"
	"sigs.k8s.io/yaml"
)

func ParseConfigs(file string) (c Config, err error) {
	conf, _ := os.ReadFile(file)
	yaml.Unmarshal(conf, &c)

	if err := c.Validate(); err != nil {
		return Config{}, errors.New("Configuration is not valid - "+err.Error())
	}

	return c, nil
}
