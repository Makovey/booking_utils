package env

import (
	"fmt"

	"github.com/Makovey/booking_utils/pkg/logger"
)

type AuthConfig struct {
	authDatabaseDSN string
	authPort        string
}

func NewAuthConfig(log logger.Logger) (*AuthConfig, error) {
	fn := "env.NewAuthConfig"

	aggregator, err := newEnvAggregator(log)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	log.Debug("AuthDatabaseDSN: " + aggregator.authDatabaseDSN)
	log.Debug("AuthPort: " + aggregator.authPort)

	return &AuthConfig{
		authDatabaseDSN: aggregator.authDatabaseDSN,
		authPort:        aggregator.authPort,
	}, nil
}

func (c *AuthConfig) AuthDatabaseDSN() string {
	return c.authDatabaseDSN
}

func (c *AuthConfig) AuthPort() string {
	return c.authPort
}
