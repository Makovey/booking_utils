package config

type Config interface {
	AuthDatabaseDSN() string
	AuthPort() string
}
