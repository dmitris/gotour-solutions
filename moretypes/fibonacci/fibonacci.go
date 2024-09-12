package fibonacci

// https://go.dev/tour/moretypes/26
// Implement a fibonacci function that returns a function (a closure) that
// returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	f0 := 0
	f1 := 1
	// cnt := 0
	return func() int {
		// fmt.Printf("cnt: %d, f0: %d, f1: %d\n", cnt, f0, f1)
		// cnt++
		ret := f0
		newf1 := f0 + f1
		f0 = f1
		f1 = newf1
		return ret
	}
}
