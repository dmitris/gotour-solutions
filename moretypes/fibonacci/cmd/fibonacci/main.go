package main

import (
	"fmt"

	"github.com/dmitris/gotour-solutions/moretypes/fibonacci"
)

func main() {
	f := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, f())
	}
}
