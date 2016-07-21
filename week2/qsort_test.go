package main

import (
	"testing"
)

func TestQsort(t *testing.T) {
	tests := []struct {
		file                string
		want1, want2, want3 int
	}{
		{"10.txt", 25, 29, 21},
		{"100.txt", 615, 587, 518},
		{"1000.txt", 10297, 10184, 8921},
	}

	for _, c := range tests {
		sample := readFile(c.file)
		got := qsort(append([]int{}, sample...))
		if got != c.want1 {
			t.Errorf("File %q, first pivot, want %v, got %v", c.file, c.want1, got)
		}
	}

	choosePartition = chooseLast
	for _, c := range tests {
		sample := readFile(c.file)
		got := qsort(append([]int{}, sample...))
		if got != c.want2 {
			t.Errorf("File %q, last pivot, want %v, got %v", c.file, c.want2, got)
		}
	}

	choosePartition = chooseMedn
	for _, c := range tests {
		sample := readFile(c.file)
		got := qsort(append([]int{}, sample...))
		if got != c.want3 {
			t.Errorf("File %q, median pivot, want %v, got %v", c.file, c.want3, got)
		}
	}
}
