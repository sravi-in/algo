package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	g    Graph
	want []Edge
}{
	{
		Graph{
			1: {2, 3, 4, 7},
			2: {1, 3, 4},
			3: {1, 2, 4},
			4: {1, 2, 3, 5},
			5: {4, 6, 7, 8},
			6: {5, 7, 8},
			7: {1, 5, 6, 8},
			8: {5, 6, 7},
		},
		1,2
		d1,3
		1,4
		1,7
		2,3
		2,4
		3,4
		4,5
		5,6
		5,7
		5,8
		6,7
		6,8
		7,8
		[]Edge{{1, 7}, {4, 5}},
	},
}

func TestMinCut(t *testing.T) {
	for _, c := range tests {
		got := minCut(c.g)
		if reflect.DeepEqual(c.want, got) == false {
			t.Errorf("want %v, got %v", c.want, got)
		}
	}
}
