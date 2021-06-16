package greedyalgorithms

type TreeHuffman struct {
	Map  map[string]int
	Tree [][]int
}

func (t *TreeHuffman) Insert(value []int) {
	if len(t.Tree) == 0 {
		t.Tree = append(t.Tree, []int{0, 0})
		t.Tree = append(t.Tree, value)
		return
	}
	t.SiftUp(value)
}

func (t *TreeHuffman) SiftUp(value []int) {
	t.Tree = append(t.Tree, value)
	idx := len(t.Tree) - 1
	for {
		if !(idx/2 > 0 && idx/2 <= len(t.Tree)-1) {
			break
		}
		if value[0] >= t.Tree[idx/2][0] {
			return
		}
		if value[0] < t.Tree[idx/2][0] {
			t.Tree[idx], t.Tree[idx/2] = t.Tree[idx/2], t.Tree[idx]
			idx /= 2
		}
	}
}

func (t *TreeHuffman) SiftDown() {
	t.Tree[1] = t.Tree[len(t.Tree)-1]
	t.Tree = append(t.Tree[0:1], t.Tree[1:len(t.Tree)-1]...)
	check := func(v1 int, v2 int, v3 int) (int, bool) {
		if t.Tree[v2][0] <= t.Tree[v3][0] && t.Tree[v1][0] > t.Tree[v2][0] {
			t.Tree[v2], t.Tree[v1] = t.Tree[v1], t.Tree[v2]
			return v2, false
		} else if t.Tree[v2][0] > t.Tree[v3][0] && t.Tree[v1][0] > t.Tree[v3][0] {
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
			idx, con = check(idx, idx*2, len(t.Tree)-1)
			if con {
				break
			}
			continue
		} else {
			break
		}
	}
}

func (t *TreeHuffman) GetMin() []int {
	min := t.Tree[1]
	t.SiftDown()
	return min
}

func freq(str string) map[string]int {
	runeStr := []rune(str)
	mapStr := make(map[string]int)

	var result []int

	for idx := range runeStr {
		if _, value := mapStr[string(runeStr[idx])]; value {
			mapStr[string(runeStr[idx])]++
		} else {
			mapStr[string(runeStr[idx])] = 1
		}
	}
	return mapStr
}

// func (t *TreeHuffman) Huffman(str string) {
// }
