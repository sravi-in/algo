package unionfind

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewSet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sizes := append([]int{0}, rand.Perm(65536)[:20]...)
	for _, n := range sizes {
		got := NewSet(n)
		if len(got) != n {
			t.Fatalf("len(NewSet(%d)) = %d, want %d", n, len(got), n)
		}
		for i, j := range got {
			if i != j {
				t.Fatalf("NewSet(%d)[%d] = %d, want %d", n, i, j, i)
			}
		}
	}
}

func TestUnionFind(t *testing.T) {
	set := NewSet(10)

	if set.Find(3) == set.Find(4) {
		t.Fatalf("Find(3) != Find(4)")
	}

	set.Union(3, 4)

	if set.Find(3) != set.Find(4) {
		t.Fatalf("Find(3) != Find(4)")
	}
}
