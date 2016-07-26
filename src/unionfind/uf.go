package unionfind

type UnionFind []int

func NewSet(n int) (uf UnionFind) {
	uf = make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return
}

func (uf UnionFind) Union(p, q int) {
	pRoot, pRank := uf.intFind(p)
	qRoot, qRank := uf.intFind(q)

	if pRank < qRank {
		uf[pRoot] = qRoot
	} else {
		uf[qRoot] = pRoot
	}
}

func (uf UnionFind) Find(p int) int {
	root, _ := uf.intFind(p)
	return root
}

func (uf UnionFind) intFind(p int) (int, int) {
	i := 0
	for p--; uf[p] != p; p = uf[p] {
		i++
	}
	return p, i
}
