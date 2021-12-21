package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig Load config with viper
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}
}
