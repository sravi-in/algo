package main

import (
	"fmt"
)

const (
	MAX_VERTICES = 875715
)

type Graph []Siblings
type Siblings []int

func main() {
	g := readGraphRev("SCC.txt")
	fmt.Println(len(g))
}

func readGraphRev(file string) Graph {
}
