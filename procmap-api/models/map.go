package models

type Map struct {
	ID      uint    `json:"id"`
	Seed    uint    `json:"seed"`
	Alpha   float32 `json:"alpha"`
	Octaves uint    `json:"octaves"`
}

// var (
// 	seed    = 1.0
// 	alpha   = 1.08
// 	beta    = 0.0
// 	octaves = 6
// )
