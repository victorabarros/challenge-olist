package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config summarises environment variables.
type Config struct {
	Database Database `mapstructure:",squash"`
	Server   Server   `mapstructure:",squash"`
	LogLevel string   `mapstructure:"log_level"`
}

// Database summarises Database variables.
type Database struct {
	Driver   string `mapstructure:"database_driver"`
	Host     string `mapstructure:"database_host"`
	Name     string `mapstructure:"database_name"`
	Password string `mapstructure:"database_password"`
	User     string `mapstructure:"database_user"`

	DSN string
}

// Server summarises Server variables.
type Server struct {
	Port int `mapstructure:"port"`
}

var (
	cfg *Config
)

// Load return environment variables loaded.
func Load() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}
	fmt.Println("Loading envs.")

	viper.SetDefault("DATABASE_DRIVER", "mysql")
	viper.SetDefault("DATABASE_HOST", "127.0.0.1:8093")
	viper.SetDefault("DATABASE_NAME", "olist")
	viper.SetDefault("DATABASE_PASSWORD", "1234")
	viper.SetDefault("DATABASE_USER", "olist")
	viper.SetDefault("LOG_LEVEL", "INFO")
	viper.SetDefault("PORT", 8092)

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.Database.DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Name)

	return cfg, nil
}
