package controller

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
    return c.SendString("get all cards")
}

func GetSpecific(c *fiber.Ctx) error {
    return c.SendString("specific card")
}

func NewCard(c *fiber.Ctx) error {
    return c.SendString("new card")
}

func UpdateCard(c *fiber.Ctx) error {
    return c.SendString("update card")
}
