package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	EnabledTests    []string `json:"enabled_tests"`
	IntervalSeconds int      `json:"interval_seconds,omitempty"`
}

func Parse(data []byte) (*Config, error) {
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("config: invalid JSON: %w", err)
	}
	return &c, nil
}