package procgen

import (
	"github.com/ojrac/opensimplex-go"
)

type NoiseMap struct {
	seed      int64
	noise     opensimplex.Noise
	exponent  float64
	frequency float64
}

func NewNoiseMap(seed int64, exponent float64, frequency float64) *NoiseMap {
	return &NoiseMap{
		seed:      seed,
		noise:     opensimplex.NewNormalized(seed),
		exponent:  exponent,
		frequency: frequency,
	}
}

func (n *NoiseMap) GetNoiseMap(x, y int) float64 {
	freq := n.frequency
	xNoise := float64(x) * freq
	yNoise := float64(y) * freq

	ret := n.noise.Eval2(xNoise, yNoise)
	return ret
}
