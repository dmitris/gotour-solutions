package main

import (
	"fmt"

	"github.com/dmitris/gotour-solutions/moretypes/fibonacci"
)

func main() {
	f := fibonacci.Fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Printf("%2d: %3d\n", i, f())
	}
}
