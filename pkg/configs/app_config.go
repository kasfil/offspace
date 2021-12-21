package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// FiberConfig Fiber app configuration
func FiberConfig() fiber.Config {
  return fiber.Config{
    Prefork:               viper.Get("stage") == "production",
    DisableStartupMessage: viper.Get("stage") == "production",
    AppName:               viper.GetString("app_name"),
    ReadTimeout:           time.Second * time.Duration(viper.GetInt("read_timeout")),
  }
}
