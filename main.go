package main

import (
	"offspace/pkg/configs"
	"offspace/pkg/routes"
	"offspace/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// load config
	configs.LoadConfig()

	// connect and migrate database
	db := utils.DBConnect()
	utils.RunMigration(db)

	bootstrap := fiber.New(configs.FiberConfig())

	routes.PublicRoutes(bootstrap)

	bootstrap.Listen(":8000")
}
