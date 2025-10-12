// config/config.go
package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

//corresponding to .env
type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET_KEY"`
}

// DSN:database source name
func (c *Config) DSN() string {
	// "postgres://user:password@host:port/dbname?sslmode=disable"
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}

// LoadConfig:load .env file
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")        // find file dir 
	viper.SetConfigName(".env")     // find ".env"
	viper.SetConfigType("env")      // tell file struct "env"

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed load .env file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Failed analyze .env file: %w", err)
	}

	return &config, nil
}
