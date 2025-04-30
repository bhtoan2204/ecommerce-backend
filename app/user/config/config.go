package config

import (
	"fmt"
	"os"
	"strings"
	"user/package/settings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitLoadConfig() (settings.Config, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	if env != "production" {
		if err := godotenv.Load(
			fmt.Sprintf(".env.%s", env),
		); err != nil {
			panic(fmt.Errorf("error loading .env files: %w", err))
		}
	}

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnv(v)

	var config settings.Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}

	return config, nil
}

func bindEnv(v *viper.Viper) {
	// Server mappings
	v.BindEnv("server.mode", "SERVER_MODE")
	v.BindEnv("server.gin_mode", "GIN_MODE")

	// Security mappings
	v.BindEnv("security.jwt_access_secret", "SECURITY_JWT_ACCESS_SECRET")
	v.BindEnv("security.jwt_refresh_secret", "SECURITY_JWT_REFRESH_SECRET")
	v.BindEnv("security.jwt_access_expiration", "SECURITY_JWT_ACCESS_EXPIRATION")
	v.BindEnv("security.jwt_refresh_expiration", "SECURITY_JWT_REFRESH_EXPIRATION")

	// // Kafka mappings
	// v.BindEnv("kafka.broker", "KAFKA_BROKER")
	// v.BindEnv("kafka.port", "KAFKA_PORT")
	// v.BindEnv("kafka.topic", "KAFKA_TOPIC")
	// v.BindEnv("kafka.group_id", "KAFKA_GROUP_ID")

	// // Opentelemetry mappings
	// v.BindEnv("opentelemetry.endpoint", "OPENTELEMETRY_ENDPOINT")
}
