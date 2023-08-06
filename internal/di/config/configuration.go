package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFile = "config.json"

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func LoadConfig() (*Configuration, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("open config file %s: %w", configFile, err)
	}

	var config Configuration
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode congig file %s: %w", configFile, err)
	}

	return &config, nil
}
