package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	LogLevel string
	Port  string

	CommentServiceHost string
	CommentServicePort int

	UserServiceHost string
	UserServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(GetFromOs("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(GetFromOs("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(GetFromOs("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(GetFromOs("POSTGRES_DATABASE", "insta_post"))
	c.PostgresUser = cast.ToString(GetFromOs("POSTGRES_USER", "khusniddin"))
	c.PostgresPassword = cast.ToString(GetFromOs("POSTGRES_PASSWORD", "1234"))
	c.LogLevel = cast.ToString(GetFromOs("LOG_LEVEL", "debug"))

	c.Port = cast.ToString(GetFromOs("PORT", ":9000"))

	c.CommentServiceHost = cast.ToString(GetFromOs("COMMENT_SERVICE_HOST", "localhost"))
	c.CommentServicePort = cast.ToInt(GetFromOs("COMMENT_SERVICE_PORT", 9001))

	c.UserServiceHost = cast.ToString(GetFromOs("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(GetFromOs("USER_SERVICE_PORT", 9002))

	return c
}

// Get from os (if Exists)
func GetFromOs(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
