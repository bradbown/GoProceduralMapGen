package models

type Map struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	UserID   uint    `json:"user_id" gorm:"foreignKey"`
	Name     string  `json:"name"`
	Seed     uint    `json:"seed"`
	Alpha    float32 `json:"alpha"`
	Octaves  uint    `json:"octaves"`
	NoiseMap string  `json:"noise_map"`
}

// var (
// 	seed    = 1.0
// 	alpha   = 1.08
// 	beta    = 0.0
// 	octaves = 6
// )
