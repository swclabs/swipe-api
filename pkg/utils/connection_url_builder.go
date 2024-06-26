package utils

import (
	"fmt"

	"swclabs/swipecore/internal/config"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.DbHost,
			config.DbPort,
			config.DbUser,
			config.DbPassword,
			config.DbName,
			config.DbSSLMode,
		)
	case "pg-migrate":
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			config.DbPort,
			config.DbName,
			config.DbSSLMode,
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			config.DbPort,
			config.DbName,
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			config.RedisHost,
			config.RedisPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}

func ConnectionURLBuilderWithEnv(n string, env config.Env) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			env.DbHost,
			env.DbPort,
			env.DbUser,
			env.DbPassword,
			env.DbName,
			env.DbSSLMode,
		)
	case "pg-migrate":
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			env.DbUser,
			env.DbPassword,
			env.DbHost,
			env.DbPort,
			env.DbName,
			env.DbSSLMode,
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			env.DbUser,
			env.DbPassword,
			env.DbHost,
			env.DbPort,
			env.DbName,
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			env.RedisHost,
			env.RedisPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
