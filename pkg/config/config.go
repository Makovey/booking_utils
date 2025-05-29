package config

type AuthConfig interface {
	AuthDatabaseDSN() string
	AuthPort() string
}
