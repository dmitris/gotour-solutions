package main

import (
	"slices"
	"testing"

	"golang.org/x/tour/tree"
)

func TestCollectChannel(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
	want := []int{0, 1, 2, 3, 4}
	s := collectChannel(ch)
	if !slices.Equal(s, want) {
		t.Errorf("wrong result %#v, want %#v", s, want)
	}
}

func TestWalk(t *testing.T) {
	tests := []struct {
		name   string
		t1, t2 *tree.Tree
		want   bool
	}{
		{
			"nil",
			nil, nil,
			true,
		},
		{
			"left nil right not",
			nil, tree.New(2),
			false,
		},
		{
			"same number 7 ",
			tree.New(7), tree.New(7),
			true,
		},
		{
			"same number 12",
			tree.New(12), tree.New(12),
			true,
		},
		{
			"different trees",
			tree.New(3), tree.New(4),
			false,
		},
	}

	for _, tt := range tests {
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		go Walk(tt.t1, ch1)
		go Walk(tt.t2, ch2)
		s1 := collectChannel(ch1)
		s2 := collectChannel(ch2)
		if res := slices.Equal(s1, s2); res != tt.want {
			t.Errorf("%s - got %t, want %t, t1: %s, t2: %s, s1: %v, s2: %v",
				tt.name, res, tt.want, tt.t1, tt.t2, s1, s2)
		}
	}
}
func TestSame(t *testing.T) {
	tests := []struct {
		name   string
		t1, t2 *tree.Tree
		want   bool
	}{
		{
			"nil",
			nil, nil,
			true,
		},
		{
			"same trees",
			tree.New(7), tree.New(7),
			true,
		},
		{
			"different trees",
			tree.New(4), tree.New(3),
			false,
		},
	}
	for _, tt := range tests {
		if res := Same(tt.t1, tt.t2); res != tt.want {
			t.Errorf("%s got %t, want %t, t1: %s, t2: %s",
				tt.name, res, tt.want, tt.t1, tt.t2)
		}
	}
}
