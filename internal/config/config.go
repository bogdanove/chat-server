package config

import "github.com/joho/godotenv"

// Load - load environments
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig - configuration for GRPS server
type GRPCConfig interface {
	Address() string
}

// PGConfig - configuration for PG database
type PGConfig interface {
	DSN() string
}
