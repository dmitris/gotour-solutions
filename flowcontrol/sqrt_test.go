package flowcontrol

import (
	"math"
	"math/rand/v2"
	"testing"
)

func TestSqrt(t *testing.T) {
	iterations := 100
	for range iterations {
		n := rand.IntN(math.MaxInt32)
		myRoot := Sqrt(float64(n))
		realRoot := math.Sqrt(float64(n))
		if math.Abs(myRoot-realRoot) > tolerance {
			t.Errorf("n=%d, computed root: %f, stdlib math root: %f\n",
				n, myRoot, realRoot)
		}
	}
}
