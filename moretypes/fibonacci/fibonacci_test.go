package fibonacci

import "testing"

// https://en.wikipedia.org/wiki/Fibonacci_sequence
func TestFibonacci(t *testing.T) {
	tests := []struct {
		inp, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, tt := range tests {
		f := Fibonacci()
		fib := 0
		for range tt.inp + 1 {
			fib = f()
		}
		if fib != tt.want {
			t.Errorf("input %d, got %d, want %d", tt.inp, fib, tt.want)
		}
	}
}
