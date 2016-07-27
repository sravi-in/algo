package unionfind

type UnionFind struct {
	id, sz []int
}

func NewSet(n int) *UnionFind {
	uf := UnionFind{id: make([]int, n), sz: make([]int, n)}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return &uf
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)

	if uf.sz[pRoot] < uf.sz[qRoot] {
		uf.id[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	} else {
		uf.id[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}
}

func (uf *UnionFind) Find(p int) int {
	// Internal index starts with 0
	p--
	for uf.id[p] != p {
		//Lazy path compression, set id to grandparent everytime
		uf.id[p] = uf.id[uf.id[p]]
		p = uf.id[p]
	}
	return p
}
