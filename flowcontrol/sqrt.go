package flowcontrol

import (
	"math"
)

const (
	maxIterations = 40
	tolerance     = 0.001
)

func Sqrt(x float64) float64 {
	z := x / 2
	// fmt.Printf("input %f (%[1]T)\n", x)
	for i := 0; i < maxIterations && math.Abs(z*z-x) > tolerance; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("i: %d, z: %f, z*z: %f, diff: %f\n", i, z, z*z, math.Abs(z*z-x))
	}
	return z
}
