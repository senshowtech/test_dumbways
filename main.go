package main

import (
	controller "dumbways/controller"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const jwtSecret = "secret"

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/consume/transaction", controller.ConsumeTransaction)
	app.Post("/create/transaction", controller.PublishTransaction)

	app.Post("/login", controller.Login)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
