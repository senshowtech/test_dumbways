package main

import (
	controller "dumbways/controller"
	middlewares "dumbways/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const jwtSecret = "secret"

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/consume/wallet", controller.ConsumeTransaction)
	app.Post("/create/wallet", controller.PublishTransaction)
	app.Get("/wallet", middlewares.AuthRequired(), controller.Transaction)

	app.Post("/login", controller.Login)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
