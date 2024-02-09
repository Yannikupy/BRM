package factory

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func SetConfigs() error {
	viper.SetConfigFile("config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read config file %w", err)
	}
	if err := godotenv.Load("config/.env"); err != nil {
		return fmt.Errorf("unable to load .env file %w", err)
	}
	return nil
}
