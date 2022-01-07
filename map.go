// package main

// import (
// 	"fmt"
// 	"math"

// 	"github.com/iand/perlin"
// )

// var (
// 	seed    = 1.0
// 	alpha   = 1.08
// 	beta    = 0.0
// 	octaves = 6
// )

// type Map struct {
// 	width, height int
// 	scale         float32
// }

// func GenerateNoiseMap() {

// 	mapGen := Map{
// 		width:  10,
// 		height: 10,
// 		scale:  2,
// 	}

// 	//var noiseMap [10][10]float64

// 	for x := 0; x < mapGen.width; x++ {
// 		for y := 0; y < mapGen.height; y++ {

// 			perlinValue := perlin.Noise2D(float64(x), float64(y), int64(seed), alpha, beta, octaves)
// 			perlinValue = (perlinValue * .5) + .5
// 			perlinValue = math.Max(0.0, perlinValue)
// 			perlinValue = math.Min(1.0, perlinValue)

// 			//noiseMap[x][y] = perlinValue
// 			fmt.Println(perlinValue)
// 		}
// 	}
// }
