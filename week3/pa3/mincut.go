package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unionfind"
)

type Graph struct {
	V map[int]int
	E []Edge
}

type Edge struct {
	u, v int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	g := readGraph("kargerMinCut.txt")
	fmt.Println(minCut(g))
}

func minCut(g *Graph) (cuts []Edge) {
	n := len(g.V)
	cuts = randomContraction(g)
	// Run Krager's random contraction for n^2.log(n) trials
	// and remember mincut in cuts
	for t := n * n * int(math.Log10(float64(n))); t > 0; t-- {
		k := randomContraction(g)
		if len(k) < len(cuts) {
			cuts = k
		}
	}
	return
}

func randomContraction(g *Graph) (cuts []Edge) {
	n := len(g.V)
	uf := unionfind.NewSet(n)
	for _, i := range rand.Perm(len(g.E)) {
		e := g.E[i]
		if n <= 2 {
			break
		}
		// Self-loop
		if uf.Find(e.u) == uf.Find(e.v) {
			continue
		}
		uf.Union(e.u, e.v)
		n--
	}

	if n > 2 {
		log.Fatal("n > 2")
	}

	for _, e := range g.E {
		// An edge not part of same set is a cut
		if uf.Find(e.u) != uf.Find(e.v) {
			cuts = append(cuts, e)
		}
	}
	return
}

func readGraph(file string) *Graph {
	g := NewGraph()
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vertices := strings.Split(scanner.Text(), "\t")
		u, _ := strconv.Atoi(vertices[0])
		for _, v := range vertices[1:] {
			v, _ := strconv.Atoi(v)
			g.AddEdge(u, v)
		}
	}
	return g
}

func NewGraph() *Graph {
	return &Graph{V: make(map[int]int)}
}

func (g *Graph) AddEdge(u, v int) {
	if u > v {
		return
	}
	g.E = append(g.E, Edge{u, v})
	g.V[u] = 1
	g.V[v] = 1
}
