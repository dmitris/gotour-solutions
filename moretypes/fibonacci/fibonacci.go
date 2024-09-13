package fibonacci

// https://go.dev/tour/moretypes/26
// Implement a fibonacci function that returns a function (a closure) that
// returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// Fibonacci is a function that returns a function that returns
// the next fibonacci number - // https://en.wikipedia.org/wiki/Fibonacci_sequence.
func Fibonacci() func() int {
	f0, f1 := 0, 1
	// cnt := 0
	return func() int {
		// fmt.Printf("cnt: %d, f0: %d, f1: %d\n", cnt, f0, f1)
		// cnt++
		ret := f0
		f0, f1 = f1, f0+f1
		return ret
	}
}
