package main

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch. It closes
// the channel ch on the function exit.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

// walk recursively walks the tree and sends values
// to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees t1 and t2
// contain the same values. It calls Walk,
// collects the values from the channels filled by
// Walk into slilces and compares the slices for equality.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	s1 := collectChannel(ch1)
	s2 := collectChannel(ch2)
	return slices.Equal(s1, s2)
}
func main() {
	t1, t2 := tree.New(5), tree.New(5)
	t3, t4 := tree.New(3), tree.New(4)
	cleaned := removeParensAll(t1, t2, t3, t4)

	fmt.Printf("t1: %s =>\n%s\nt2 : %s =>\n%s\n",
		t1, cleaned[0], t2, cleaned[1])
	isSame := Same(t1, t2)
	fmt.Printf("Same(t1, t2): %t\n\n", isSame)
	fmt.Printf("t3: %s =>\n%s\nt4: %s =>\n%s\n",
		t3, cleaned[2], t4, cleaned[3])
	isSame = Same(t3, t4)
	fmt.Printf("Same(t3, t4): %t\n", isSame)
}

func removeParens(t *tree.Tree) string {
	s := t.String()
	s = strings.ReplaceAll(s, "(", "")
	return strings.ReplaceAll(s, ")", "")
}
func removeParensAll(trees ...*tree.Tree) []string {
	ret := make([]string, 0, len(trees))
	for _, t := range trees {
		ret = append(ret, removeParens(t))
	}
	return ret
}

func collectChannel(ch <-chan int) []int {
	s := []int{}
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// // generic version:
// func collectChannel[T any](ch <-chan T) []T {
// 	s := []T{}
// 	for v := range ch {
// 		s = append(s, v)
// 	}
// 	return s
// }
