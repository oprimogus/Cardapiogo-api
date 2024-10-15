package config

type Environment string

const (
	Production Environment = "production"
	Staging    Environment = "staging"
	Local      Environment = "local"
)
