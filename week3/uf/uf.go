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
	pRoot := uf.Find(p - 1)
	qRoot := uf.Find(q - 1)

	uf[pRoot] = qRoot
}

func (uf UnionFind) Find(p int) int {
	for p--; uf[p] != p; p = uf[p] {
	}
	return p + 1
}
