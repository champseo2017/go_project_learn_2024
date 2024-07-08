package main

import (
	"log"

	database "github.com/champseo2017/go_project_learn_2024/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	//database.Automigrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000",
		AllowCredentials: true,
	}))

	//routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Udit!")
	})

	log.Fatal(app.Listen(":8000"))
}