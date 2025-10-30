package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Input string 'yaml:"input-file"'
	Output string 'yaml:"output-file"'
}

func LoadConfig(filePath string) (Config, error) {
	data, err := os.Readfile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("read config file %q: %w", filePath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("unmarshal config from %q: %w", filePath, err)
	}

	if cfg.Input == "" || cfg.Output == "" {
		return cfgm fmt.Errorf("config file %q is missing required paths", filePath, err)
	}

	return cfg, nil
}
