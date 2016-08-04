package dheap

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestHeap(t *testing.T) {
	dh := NewDHeap()
	for v, score := range rand.Perm(100000) {
		dh.Insert(&DScore{V: v, Score: score})
	}

	for i := 0; i < 100000; i++ {
		d := dh.Remove()
		testHeap(t, dh, d, i)
	}
}

func testHeap(t *testing.T, dh *DHeap, d *DScore, i int) {
	if d.Score != i {
		t.Fatalf("%dth remove returned %d, expected %d %v\n", i, d.Score, i, dh)
	}

	if len(dh.m) != len(dh.h) {
		t.Fatalf("%dth remove len(dh.m) %d != len(dh.h) %d\n",
			i, len(dh.m), len(dh.h))
	}

	if _, ok := dh.m[d.V]; ok {
		t.Fatalf("%dth remove entry present in map\n", i)
	}
}
