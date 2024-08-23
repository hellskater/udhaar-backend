package cmd

import (
	"fmt"

	"github.com/hellskater/udhaar-backend/internal/router"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DevMode  bool   `mapstructure:"dev" yaml:"dev"`
	Origin   string `mapstructure:"origin" yaml:"origin"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Postgres struct {
		User         string `mapstructure:"user" yaml:"user"`
		Password     string `mapstructure:"password" yaml:"password"`
		Host         string `mapstructure:"host" yaml:"host"`
		Port         string `mapstructure:"port" yaml:"port"`
		DatabaseName string `mapstructure:"databaseName" yaml:"databaseName"`
	} `mapstructure:"postgres" yaml:"postgres"`
}

func init() {
	viper.SetDefault("dev", true)
	viper.SetDefault("origin", "http://localhost:8080")
	viper.SetDefault("port", 8080)
	viper.SetDefault("postgres.user", "postgres")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.databaseName", "udhaar")
}

func (c Config) getDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.DatabaseName)
	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if config.DevMode {
		engine.Logger.LogMode(logger.Info)
	}

	return engine, nil
}

func provideRouterConfig(c *Config) *router.Config {
	return &router.Config{
		Origin:      c.Origin,
		Development: c.DevMode,
	}
}