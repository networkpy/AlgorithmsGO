package greedyalgorithms

type Tree struct {
	Tree []int
}

func (t *Tree) Insert(value int) {
	if len(t.Tree) == 0 {
		t.Tree = append(t.Tree, 0)
		t.Tree = append(t.Tree, value)
		return
	}
	t.SiftUp(value)
}

func (t *Tree) SiftUp(value int) {
	t.Tree = append(t.Tree, value)
	idx := len(t.Tree) - 1
	for {
		if !(idx/2 > 0 && idx/2 <= len(t.Tree)-1) {
			break
		}
		if value <= t.Tree[idx/2] {
			return
		}
		if value > t.Tree[idx/2] {
			t.Tree[idx], t.Tree[idx/2] = t.Tree[idx/2], t.Tree[idx]
			idx /= 2
		}
	}
}

func (t *Tree) SiftDown() {
	t.Tree[1] = t.Tree[len(t.Tree)-1]
	t.Tree = append(t.Tree[0:1], t.Tree[1:len(t.Tree)-1]...)
	check := func(v1 int, v2 int, v3 int) (int, bool) {
		if t.Tree[v2] >= t.Tree[v3] && t.Tree[v1] < t.Tree[v2] {
			t.Tree[v2], t.Tree[v1] = t.Tree[v1], t.Tree[v2]
			return v2, false
		} else if t.Tree[v2] < t.Tree[v3] && t.Tree[v1] < t.Tree[v3] {
			t.Tree[v3], t.Tree[v1] = t.Tree[v1], t.Tree[v3]
			return v3, false
		} else {
			return 0, true
		}
	}

	idx := 1
	con := false
	for {
		if (idx*2)+1 < len(t.Tree) {
			idx, con = check(idx, idx*2, idx*2+1)
			if con {
				break
			}
			continue
		} else if (idx * 2) < len(t.Tree) {
			idx, con = check(idx, idx*2, 0)
			if con {
				break
			}
			continue
		} else {
			break
		}
	}
}

func (t *Tree) GetMax() int {
	max := t.Tree[1]
	t.SiftDown()
	return max
}
