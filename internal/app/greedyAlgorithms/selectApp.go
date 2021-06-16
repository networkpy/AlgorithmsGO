package greedyalgorithms

import (
	"math/rand"
)

// Вывести множество отрезков, которые попарно не пересекаются
func SelApp(arr [][2]int) [][2]int {
	SortSec(arr)
	result := [][2]int{arr[0]}

	i := 1
	for i < len(arr) {
		if arr[i][0] >= result[len(result)-1][1] {
			result = append(result, arr[i])
		}
		i++
		for i < len(arr) && arr[i][0] < result[len(result)-1][1] {
			i++
		}
	}
	return result
}

// Generation sections
func GenSec(N int) [][2]int {
	var res [][2]int
	for i := 0; i < N; i++ {
		max := rand.Intn(1000)
		for max <= 0 {
			max = rand.Intn(1000)
		}
		min := rand.Intn(max)
		res = append(res, [2]int{min, max})
	}
	return res
}

func SortSec(sections [][2]int) {
	for i := 0; i < len(sections); i++ {
		for j := i + 1; j < len(sections); j++ {
			if sections[i][1] > sections[j][1] {
				sections[i], sections[j] = sections[j], sections[i]
			}
		}

	}
}
