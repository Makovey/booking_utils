package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/Makovey/booking_utils/pkg/logger"
)

const (
	authDatabaseDsnEnv = "AUTH_DATABASE_DSN"
	authPortEnv        = "AUTH_PORT"

	defaultAuthPort = "9091"
)

type resolver struct {
	log logger.Logger
}

type envAggregator struct {
	authDatabaseDSN string
	authPort        string
}

func newEnvAggregator(log logger.Logger) (*envAggregator, error) {
	fn := "env.newEnvConfig"

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("[%s]: %s", fn, err.Error())
	}

	r := &resolver{log: log}

	authDatabaseDSN, err := r.resolveVariable(authDatabaseDsnEnv, "")
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	authPort, err := r.resolveVariable(authPortEnv, defaultAuthPort)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return &envAggregator{
		authDatabaseDSN: authDatabaseDSN,
		authPort:        authPort,
	}, nil
}

func (r *resolver) resolveVariable(name, defaultValue string) (string, error) {
	fn := "env.resolveVariable"

	value := os.Getenv(name)
	if value != "" {
		return value, nil
	}

	if defaultValue != "" {
		r.log.Warnf("[%s]: %s variable is empty, but default value is %s", fn, authPortEnv, defaultAuthPort)
		return defaultValue, nil
	}

	return "", fmt.Errorf("[%s]: required variable - %s is empty", fn, name)
}
