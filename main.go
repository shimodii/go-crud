package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/go-crud/controller"
	"github.com/shimodii/go-crud/repository"
)

func init(){
    repository.OpenDatabase()
}

func main(){

    /* var app = fiber.New(fiber.Config{
        ServerHeader: "go-crud",
        AppName: "go crud application",
    }) */

    app := fiber.New()
    
    app.Get("/", controller.GetAll)
    
    app.Listen(":3000")
}
