package models

type Map struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id" gorm:"foreignKey"`
	Name      string  `json:"name"`
	Size      int     `json:"size"`
	Seed      int64   `json:"seed"`
	Exponent  float64 `json:"exponent"`
	Frequency float64 `json:"frequency"`
	NoiseMap  string  `json:"noise_map"`
}

// var (
// 	seed    = 1.0
// 	alpha   = 1.08
// 	beta    = 0.0
// 	octaves = 6
// )
