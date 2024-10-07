package main

import (
	"fmt"
)

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					break
				}
				ch <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
	}()
	return ch
}

func produce(ch1, ch2 chan<- int) {
	for i := range 11 {
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	close(ch1)
	close(ch2)
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch := merge(ch1, ch2)

	go produce(ch1, ch2)

	for v, ok := <-ch; ok; v, ok = <-ch {
		fmt.Printf("%2d\n", v)
	}
}
