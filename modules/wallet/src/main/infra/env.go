package infra

import (
	"fmt"
	"os"
)

type Env struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	KafkaServer      string
}

func NewEnv() *Env {
	return &Env{
		DatabaseHost:     getEnvValue("DATABASE_HOST", "localhost"),
		DatabasePort:     getEnvValue("DATABASE_PORT", "3306"),
		DatabaseName:     getEnvValue("DATABASE_NAME", "wallet"),
		DatabaseUsername: getEnvValue("DATABASE_USERNAME", "root"),
		DatabasePassword: getEnvValue("DATABASE_PASSWORD", "root"),
		KafkaServer:      getEnvValue("KAFKA_SERVER", "localhost"),
	}
}

func getEnvValue(name, defaultValue string) string {
	env := os.Getenv(name)
	if env != "" {
		return env
	}
	if defaultValue != "" {
		return defaultValue
	}
	panic(fmt.Errorf("%s must be set", name))
}
