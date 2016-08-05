package main

import (
	"bufio"
	"dheap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph []Edges
type Edges []EdgeHead
type EdgeHead struct {
	v, l int
}

func main() {
	g := readGraph("dijkstraData.txt")
	s := dShortPath(g, 1)
	for _, i := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		fmt.Printf("%d,", s[i])
	}
}

func dShortPath(g Graph, start int) []int {
	distances := make([]int, len(g))
	explored := make([]bool, len(g))
	dh := dheap.NewDHeap()
	for i := 1; i < len(g); i++ {
		dh.Insert(&dheap.DScore{V: i, Score: 1000000})
	}
	dh.Insert(&dheap.DScore{V: start, Score: 0})
	for d := dh.Remove(); d != nil; d = dh.Remove() {
		i, score := d.V, d.Score
		distances[i] = score
		explored[i] = true
		for _, e := range g[i] {
			if !explored[e.v] {
				dh.Insert(&dheap.DScore{V: e.v, Score: score + e.l})
			}
		}
	}
	return distances
}

// Hacky
func readGraph(file string) Graph {
	// Ignore 0 index
	g := make(Graph, 1)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \t")
		words := strings.Split(line, "\t")
		v, _ := strconv.Atoi(words[0])
		g = append(g, nil)
		for _, next := range words[1:] {
			fields := strings.Split(next, ",")
			i, _ := strconv.Atoi(fields[0])
			score, _ := strconv.Atoi(fields[1])
			g[v] = append(g[v], EdgeHead{i, score})
		}
	}
	return g
}
