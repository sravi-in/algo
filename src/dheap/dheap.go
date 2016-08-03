// package dheap is a heap used by dijkstra's algo
package dheap

type DHeap struct {
	h []*DScore
	m map[int]*DScore
}

type DScore struct {
	V, Score, index int
}

func NewDHeap() *DHeap {
	return &DHeap{m: make(map[int]*DScore)}
}

func (dh *DHeap) Insert(d *DScore) {
	val, ok := dh.m[d.V]
	// An entry already exists for v, check if the new score is
	// minimum and add
	if ok {
		if val.Score < d.Score {
			return
		}
		val.Score = d.Score
		dh.fix(val)
	} else {
		d.index = len(dh.h)
		dh.h = append(dh.h, d)
		dh.m[d.V] = d
		dh.fix(d)
	}
}

func (dh *DHeap) Remove() *DScore {
}

func (dh *DHeap) fix(d *DScore) {
	n := len(dh.h)
	i := d.index
	j := (i - 1) / 2
	for i > 0 && dh.h[i].Score < dh.h[j].Score {
		dh.h[i], dh.h[j] = dh.h[j], dh.h[i]
		dh.h[j].index, dh.h[i].index = j, i
		i = j
		j = (i - 1) / 2
	}

	for {
		i = d.index
		c1, c2 := i*2+1, i*2+2

		if c1 < n && dh.h[i] > dh.h[c1] {
			dh.h[i], dh.h[c1] = dh.h[c1], dh.h[i]
			dh.h[i].index, dh.h[c1].index = i, c1
			continue
		}

	}
}
