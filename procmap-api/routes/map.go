package routes

type Map struct {
	ID      uint    `json:"id"`
	Seed    uint    `json:"seed"`
	Alpha   float32 `json:"alpha"`
	Octaves uint    `json:"octaves"`
}
