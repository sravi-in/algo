// package dheap is a heap used by dijkstra's algo
package dheap

import "fmt"

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
		if val.Score <= d.Score {
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
	n := len(dh.h)

	if n == 0 {
		return nil
	}
	d := dh.h[0]
	dh.h[0], dh.h[n-1] = dh.h[n-1], dh.h[0]
	dh.h[0].index, dh.h[n-1].index = 0, n-1
	delete(dh.m, d.V)
	dh.h = dh.h[:n-1]
	if n > 1 {
		dh.fix(dh.h[0])
	}
	return d
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
		if c2 < n && dh.h[c2].Score < dh.h[c1].Score {
			j = c2
		} else {
			j = c1
		}

		if j < n && dh.h[i].Score > dh.h[j].Score {
			dh.h[i], dh.h[j] = dh.h[j], dh.h[i]
			dh.h[i].index, dh.h[j].index = i, j
		} else {
			break
		}

	}
}

func (dh *DHeap) String() string {
	return fmt.Sprintf("Heap: %d elements\n%v\nMap: %d elements\n%v\n", len(dh.h), dh.h, len(dh.m), dh.m)
}

func (d *DScore) String() string {
	return fmt.Sprintf("<v:%d,score:%d>", d.V, d.Score)
}
