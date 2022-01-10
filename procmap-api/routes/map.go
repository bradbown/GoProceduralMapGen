package routes

import (
	"errors"

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
		Alpha: mapM.Alpha, Octaves: mapM.Octaves}
}

func CreateNoiseMap(c *fiber.Ctx) error {
	var noiseMap models.Map

	if err := c.BodyParser(&noiseMap); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	noiseMap.NoiseMap = GenerateNoiseMap()

	database.Database.Db.Create(&noiseMap)
	responseMap := CreateResponseMap(noiseMap)

	return c.Status(200).JSON(responseMap)
}

func GetNoiseMapsFromUser(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")
	noiseMaps := []models.Map{}

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	database.Database.Db.Find(&noiseMaps, "user_id = ?", user_id)

	responseMaps := []Map{}

	for _, noiseMap := range noiseMaps {
		responseMap := CreateResponseMap(noiseMap)
		responseMaps = append(responseMaps, responseMap)
	}

	return c.Status(200).JSON(responseMaps)
}

func DeleteNoiseMap(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var noiseMap models.Map

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindNoiseMap(id, &noiseMap); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&noiseMap).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted noise map")
}

func FindNoiseMap(id int, noiseMap *models.Map) error {
	database.Database.Db.Find(&noiseMap, "id = ?", id)

	if noiseMap.ID == 0 {
		return errors.New("noise map does not exist")
	}
	return nil
}

func GenerateNoiseMap() string {

	//Todo create noise map from values
	return "test"
}
