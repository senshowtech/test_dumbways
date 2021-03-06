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

	app.Get("/consume/transaction", controller.ConsumeTransaction)
	app.Post("/create/transaction", middlewares.AuthRequired(), controller.PublishTransaction)
	app.Get("/transaction", middlewares.AuthRequired(), controller.Transaction)

	app.Get("/consume/balance", controller.ConsumeBalance)
	app.Post("/create/balance", middlewares.AuthRequired(), controller.PublishBalance)
	app.Get("/balance", middlewares.AuthRequired(), controller.Balance)

	app.Post("/login", controller.Login)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
