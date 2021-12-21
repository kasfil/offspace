package controllers

import "github.com/gofiber/fiber/v2"

// UserSignUp user sign up controllers
func UserSignUp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "hello world!"})
}
