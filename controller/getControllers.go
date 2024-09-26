package controller

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
    return c.SendString("hello from controller")
}
