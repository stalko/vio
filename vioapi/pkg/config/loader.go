package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Load from environment variables into provided Config struct and validate it using the validate struct tags.
// For more information about the validator, refer to github.com/go-playground/validator.
func Load[Config any]() (cfg Config, err error) {
	cfgPtr := new(Config)

	_ = godotenv.Load(".env")

	err = env.Parse(cfgPtr)
	if err != nil {
		return cfg, fmt.Errorf("failed to parse env: %w", err)
	}

	validate := validator.New()

	err = validate.Struct(cfgPtr)
	if err != nil {
		return cfg, fmt.Errorf("failed to validate config: %w", err)
	}

	return *cfgPtr, nil
}
