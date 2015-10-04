package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	configPath string = "config.json"
)

// If you want to implement another configuration engine
// you need to change annotations in Config struct
// and create struct with Parse method.
type Parser interface {
	Parse() (*Config, error)
}

type JsonConfigParser struct{}

func (parser *JsonConfigParser) Parse() (*Config, error) {
	str, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	config := new(Config)
	json.Unmarshal([]byte(str), &config)

	return config, nil
}
