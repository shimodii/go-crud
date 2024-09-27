package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/go-crud/controller"
	"github.com/shimodii/go-crud/repository"
)

func rootHandler(c *fiber.Ctx) error {
    return c.SendString("hello")
}

func main(){
    fmt.Println("hello world")
    
    fmt.Println("database init:")
    repository.OpenDatabase()

    var app = fiber.New(fiber.Config{
        ServerHeader: "go-crud",
        AppName: "go crud application",
    })
    
    //app.Get("/", rootHandler)
    app.Get("/", controller.GetAll)
    
    app.Listen(":3000")
}
