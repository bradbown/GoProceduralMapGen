package routes

import "github.com/bradbown/GoProceduralMapGen/models"

type Map struct {
	ID      uint    `json:"id"`
	Seed    uint    `json:"seed"`
	Alpha   float32 `json:"alpha"`
	Octaves uint    `json:"octaves"`
}

func CreateResponseMap(mapM models.Map) Map {
	return Map{ID: mapM.ID, Seed: mapM.Seed, Alpha: mapM.Alpha, Octaves: mapM.Octaves}
}
