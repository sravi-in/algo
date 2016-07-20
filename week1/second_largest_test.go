package main

import (
	"testing"
)

func TestSecondLargest(t *testing.T) {
	table := []struct {
		in   []int
		want int
	}{
		{[]int{5, 7, 3, 8, 9}, 8},
		{[]int{2, 1}, 1},
		{[]int{7, 3, 1, 5, 6, 8, 4, 2}, 7},
		{[]int{7, 3, 1, 7, 6, 7, 4, 2}, 6},
		{[]int{7, 3, 7, 7, 7, 7, 7, 7}, 3},
	}

	for _, c := range table {
		got := secondLargest(c.in)
		if got != c.want {
			t.Errorf("secondLargest(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
