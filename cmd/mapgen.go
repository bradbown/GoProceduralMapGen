/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/iand/perlin"

	"github.com/spf13/cobra"
)

var (
	seed    = 1
	alpha   = 1.08
	beta    = 0.0
	octaves = 6
)

type Map struct {
	width, height int
	scale         float32
}

// mapgenCmd represents the mapgen command
var mapgenCmd = &cobra.Command{
	Use:   "mapgen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mapgen called")
		noiseMap := GenerateNoiseMap()
		GenerateHeightMap(noiseMap)
	},
}

func init() {
	rootCmd.AddCommand(mapgenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mapgenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mapgenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GenerateNoiseMap() [900][900]float64 {

	mapGen := Map{
		width:  900,
		height: 900,
		scale:  0.0000025,
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("GenerateNoiseMap")

	var noiseMap [900][900]float64

	for x := 0; x < mapGen.width; x++ {
		for y := 0; y < mapGen.height; y++ {

			sampleX := float32(x) / mapGen.scale
			sampleY := float32(y) / mapGen.scale

			perlinValue := perlin.Noise2D(float64(sampleX), float64(sampleY), rand.Int63(), alpha, beta, octaves)
			perlinValue = math.Max(0.0, perlinValue)
			perlinValue = math.Min(1.0, perlinValue)

			noiseMap[x][y] = perlinValue
		}
	}
	return noiseMap
}

func GenerateHeightMap(noiseMap [900][900]float64) {
	rect := image.Rect(0, 0, 900, 900)
	img := image.NewGray16(rect)

	fmt.Println("GenerateHeightMap")
	for x := 0; x < 900; x++ {
		for y := 0; y < 900; y++ {
			val := noiseMap[x][y]
			col := color.Gray16{uint16(val * 0xffff)}
			img.Set(x, y, col)
		}
	}

	f, err := os.OpenFile("heightmap.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
