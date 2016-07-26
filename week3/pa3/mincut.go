package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Graph struct {
	n, m int
	E    []Edge
}

type Vertex int
type Edge struct {
	u, v Vertex
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	g := readGraph("kragerMinCut.txt")
	fmt.Println(minCut(g))
}

func minCut(g Graph) (cuts []Edge) {
	n := len(g)
	// Run Krager's random contraction for n^2.log(n) trials
	for t := int(n * n * math.Log(10)); t > 0; t-- {
		k := randomContraction()
	}
	return
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddEdge(u, v Vertex) {
	if u > v {
		u, v = v, u
	}
	g.E[Edge{u, v}] = 1
	g.V[u] = append(g.V[u], v)
}

func (g *Graph) ContractEdge(e Edge) {
	// delete g.E[e]
	// foreach edge in g.V[e.u]
}
