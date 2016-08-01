package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MAX_VERTICES = 875715
	GRAPH        = "SCC.txt"
	//MAX_VERTICES = 10
	//GRAPH        = "test1.txt"
	//MAX_VERTICES = 9
	//GRAPH        = "test2.txt"
	//MAX_VERTICES = 9
	//GRAPH        = "test3.txt"
	//MAX_VERTICES = 9
	//GRAPH        = "test4.txt"
	//MAX_VERTICES = 13
	//GRAPH        = "test5.txt"
)

var (
	t = 0
	s = 0
)

type Graph []Siblings
type Siblings []int
type Ftime []int
type Leader []int
type Explored []bool

type Scc struct {
	Leader, Node int
}

type Sccs []Scc

func main() {
	g := readGraphRev(GRAPH)
	var f Ftime = make([]int, MAX_VERTICES)
	var e Explored = make([]bool, MAX_VERTICES)
	finishTimes(g, f, e)
	g = readGraphFwd(GRAPH)
	e = make([]bool, MAX_VERTICES)
	var l Leader = make([]int, MAX_VERTICES)
	leader(g, f, l, e)
	sccs := getSccs(l)
	fmt.Println(len(sccs), sccs[:10])
}

func readGraphRev(file string) Graph {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var g Graph = make([]Siblings, MAX_VERTICES)
	for scanner.Scan() {
		vertices := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(vertices[0])
		for _, v := range vertices[1:] {
			v, _ := strconv.Atoi(v)
			g[v] = append(g[v], u)
		}
	}
	return g
}

func finishTimes(g Graph, f Ftime, e Explored) {
	for i := 1; i < MAX_VERTICES; i++ {
		if !e[i] {
			dfs1(g, i, f, e)
		}
	}
}

func dfs1(g Graph, start int, f Ftime, e Explored) {
	e[start] = true
	for _, j := range g[start] {
		if !e[j] {
			dfs1(g, j, f, e)
		}
	}
	t++
	f[t] = start
}

func readGraphFwd(file string) Graph {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var g Graph = make([]Siblings, MAX_VERTICES)
	for scanner.Scan() {
		vertices := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		u, _ := strconv.Atoi(vertices[0])
		for _, v := range vertices[1:] {
			v, _ := strconv.Atoi(v)
			g[u] = append(g[u], v)
		}
	}
	return g
}

func leader(g Graph, f Ftime, l Leader, e Explored) {
	for i := MAX_VERTICES - 1; i >= 1; i-- {
		v := f[i]
		if !e[v] {
			s = v
			dfs2(g, v, l, e)
		}
	}
}

func dfs2(g Graph, start int, l Leader, e Explored) {
	e[start] = true
	l[start] = s
	for _, j := range g[start] {
		if !e[j] {
			dfs2(g, j, l, e)
		}
	}
}

func getSccs(l Leader) Sccs {
	sccs := make(map[int]int)
	j := 0
	for i := 1; i < MAX_VERTICES; i++ {
		sccs[l[i]]++
		j++
	}

	i := 0
	ret := make(Sccs, len(sccs))
	for k, v := range sccs {
		ret[i] = Scc{k, v}
		i++
	}

	sort.Sort(sort.Reverse(ret))
	fmt.Println("Total:", j)
	return ret
}

func (p Sccs) Len() int           { return len(p) }
func (p Sccs) Less(i, j int) bool { return p[i].Node < p[j].Node }
func (p Sccs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
