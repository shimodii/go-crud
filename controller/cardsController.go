package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/go-crud/model"
	"github.com/shimodii/go-crud/repository"
)

func GetAll(c *fiber.Ctx) error {
    cards := []model.Card{}
    repository.Database.Db.Find(&cards)
    return c.JSON(cards)
}

func GetSpecific(c *fiber.Ctx) error {
    id := c.Params("id")
    var card model.Card
    repository.Database.Db.Find(&card, id)
    return c.JSON(card)
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
