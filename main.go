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
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello world")
	})
	app.Post("/create", controller.PublishPay)
	app.Post("/login", controller.Login)
	app.Get("/consume", controller.ConsumePay)
	app.Get("/hello", middlewares.AuthRequired(), controller.Secure)
	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
