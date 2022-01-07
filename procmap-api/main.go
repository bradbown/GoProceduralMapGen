package main

import (
	"log"

	"github.com/bradbown/GoProceduralMapGen/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}
