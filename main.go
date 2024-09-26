package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func rootHandler(c *fiber.Ctx) error {
    return c.SendString("hello")
}

func main(){
    fmt.Println("hello world")
    
    var app = fiber.New(fiber.Config{
        ServerHeader: "go-crud",
        AppName: "go crud application",
    })
    
    app.Get("/", rootHandler)
    
    app.Listen(":3000")
}
