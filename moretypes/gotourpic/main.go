package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Printf("dx: %d, dy: %d\n", dx, dy)
	dat := make([][]uint8, dx)
	for y := range dy {
		temp := make([]uint8, dx)
		for x := range dx {
			temp[x] = uint8(x ^ y)
		}
		dat[y] = temp
	}
	fmt.Printf("%#v\n", dat)
	return dat
}

// // less efficient alternative using append and slice reallocations
// func Pic2(dx, dy int) [][]uint8 {
// 	var dat [][]uint8
// 	for x := range dy {
// 		var temp []uint8
// 		for y := range dx {
// 			temp = append(temp, uint8(x^y))
// 		}
// 		dat = append(dat, temp)
// 	}
// 	return dat
// }

func main() {
	pic.Show(Pic)
}
