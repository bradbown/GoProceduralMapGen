/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math"

	"github.com/iand/perlin"

	"github.com/spf13/cobra"
)

var (
	seed    = 1.0
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
		GenerateNoiseMap()
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

func GenerateNoiseMap() {

	mapGen := Map{
		width:  10,
		height: 10,
		scale:  2,
	}

	//var noiseMap [10][10]float64

	for x := 0; x < mapGen.width; x++ {
		for y := 0; y < mapGen.height; y++ {

			perlinValue := perlin.Noise2D(float64(x), float64(y), int64(seed), alpha, beta, octaves)
			perlinValue = math.Max(0.0, perlinValue)
			perlinValue = math.Min(1.0, perlinValue)

			//noiseMap[x][y] = perlinValue
			fmt.Println(perlinValue)
		}
	}
}
