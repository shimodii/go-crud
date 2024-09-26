package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
    fmt.Println("hello world")
    var app = fiber.New(fiber.Config{
        ServerHeader: "go-crud",
        AppName: "go crud application",
    })
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("hello world from root")
    })
    app.Listen(":3000")
}
