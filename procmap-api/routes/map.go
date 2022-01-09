package routes

import (
	"github.com/bradbown/GoProceduralMapGen/database"
	"github.com/bradbown/GoProceduralMapGen/models"
	"github.com/gofiber/fiber/v2"
)

type Map struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	Name     string  `json:"name"`
	Seed     uint    `json:"seed"`
	Alpha    float32 `json:"alpha"`
	Octaves  uint    `json:"octaves"`
	NoiseMap string  `json:"noise_map"`
}

func CreateResponseMap(mapM models.Map) Map {
	return Map{ID: mapM.ID, UserID: mapM.UserID, Name: mapM.Name, Seed: mapM.Seed,
		Alpha: mapM.Alpha, Octaves: mapM.Octaves, NoiseMap: mapM.NoiseMap}
}

func CreateNoiseMap(c *fiber.Ctx) error {
	var noiseMap models.Map

	if err := c.BodyParser(&noiseMap); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&noiseMap)
	responseMap := CreateResponseMap(noiseMap)

	return c.Status(200).JSON(responseMap)
}
