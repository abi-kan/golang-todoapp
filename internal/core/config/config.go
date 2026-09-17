package core_config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	TimeZone *time.Location
}

func NewConfig() (*Config, error) {
	envVar := os.Getenv("TIME_ZONE")
	if envVar == "" {
		envVar = "UTC"
	}

	timeZone, err := time.LoadLocation(envVar)
	if err != nil {
		return nil, fmt.Errorf(
			"load time zone: %s: %w",
			envVar,
			err,
		)
	}

	config := Config{
		TimeZone: timeZone,
	}

	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get core config: %w", err)
		panic(err)
	}

	return config
}
