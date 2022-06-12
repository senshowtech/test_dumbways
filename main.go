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

	app.Get("/consume/wallet", controller.ConsumeWallet)
	app.Post("/create/wallet", controller.PublishWallet)

	app.Post("/login", controller.Login)
	app.Get("/wallet", middlewares.AuthRequired(), controller.Wallet)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
