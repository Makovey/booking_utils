package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/Makovey/booking_utils/pkg/logger"
)

type Config struct {
	authDatabaseDSN string
	authPort        string
}

func NewConfig(log logger.Logger) (*Config, error) {
	fn := "env.NewConfig"

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("[%s]: %s", fn, err.Error())
	}

	authDatabaseDSN := getValueFromEnv("AUTH_DATABASE_DSN", log)
	authPort := getValueFromEnv("AUTH_PORT", log)

	log.Debug("AuthDatabaseDSN: " + authDatabaseDSN)
	log.Debug("AuthPort: " + authPort)

	return &Config{
		authDatabaseDSN: authDatabaseDSN,
		authPort:        authPort,
	}, nil
}

func (c *Config) AuthDatabaseDSN() string {
	return c.authDatabaseDSN
}

func (c *Config) AuthPort() string {
	return c.authPort
}

func getValueFromEnv(envKey string, log logger.Logger) string {
	fn := "env.getValueFromEnv"

	value := os.Getenv(envKey)

	if value == "" {
		log.Errorf("[%s]: value from %s is empty", fn, envKey)
	}

	return value
}
