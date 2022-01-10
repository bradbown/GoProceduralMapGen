package main

import (
	"log"

	"github.com/bradbown/GoProceduralMapGen/database"
	"github.com/bradbown/GoProceduralMapGen/models"
	"github.com/bradbown/GoProceduralMapGen/routes"
	"github.com/gofiber/fiber/v2"
)

var user models.User

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.FindUser(1, &user)

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	app.Post("/api/maps", routes.CreateNoiseMap)
	app.Get("/api/maps/:user_id", routes.GetNoiseMapsFromUser)
	app.Delete("api/maps/:id", routes.DeleteNoiseMap)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API " + user.FirstName)
}
