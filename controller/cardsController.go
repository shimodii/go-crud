package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/go-crud/model"
	"github.com/shimodii/go-crud/repository"
	"github.com/shimodii/go-crud/service"
)

func GetAll(c *fiber.Ctx) error {
    cards := []model.Card{}
    repository.Database.Db.Find(&cards)
    return c.JSON(cards)
}

func GetSpecific(c *fiber.Ctx) error {
    id := c.Params("id")
    var card model.Card

    err := repository.Database.Db.Find(&card, id)
    if err != nil {
        return c.Status(404).JSON("card not found")
    }
    //if (service.ExistCheck(card)) {
    //    return c.Status(404).JSON("card not found")
    //}
    
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
    id := c.Params("id")
    
    var card model.Card

    repository.Database.Db.Find(&card, "id = ?", id)

    if (service.ExistCheck(card)) {
        c.Status(404).JSON("card not found")
    }

    type newCardData struct {
        Number string `json:"number"`
        Owner string `json:"owner"`
    }
    var newCard newCardData

    if err := c.BodyParser(&newCard); err != nil {
        c.Status(500).JSON(err.Error())
    }

    card.Owner = newCard.Owner
    card.Number = newCard.Number

    repository.Database.Db.Save(&card)
    
    return c.Status(200).JSON(card)
}

func DeleteCard(c *fiber.Ctx) error {
    id := c.Params("id")
    var card model.Card

    repository.Database.Db.Find(&card, "id = ?", id)
    
    if (service.ExistCheck(card)) {
        c.Status(404).JSON("card not found")
    }

    err := repository.Database.Db.Delete(&card).Error
    if err != nil {
        return c.SendString("failed to Delete")
    }

    return c.Status(200).JSON("successfully deleted dude")
}
