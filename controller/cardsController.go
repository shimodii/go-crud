package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/go-crud/model"
	"github.com/shimodii/go-crud/repository"
)

func GetAll(c *fiber.Ctx) error {
    return c.SendString("get all cards")
}

func GetSpecific(c *fiber.Ctx) error {
    return c.SendString("specific card")
}

func NewCard(c *fiber.Ctx) error {
    var card model.Card

    if err := c.BodyParser(&card); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    repository.Database.Db.Create(&card)
    return c.SendString("OK")
}

func UpdateCard(c *fiber.Ctx) error {
    return c.SendString("update card")
}
