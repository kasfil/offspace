package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// StaticConfig fiber static configuration
func StaticConfig() fiber.Static {
	return fiber.Static{
		Compress:      true,
		Index:         "404.html",
		CacheDuration: time.Second * time.Duration(viper.GetInt("static_cache_duration")),
	}
}
