package greedyalgorithms

import (
	"fmt"
	"sort"
)

func GreedyBackpack() {
	_, w, arr := ReadItem()
	sort.Slice(arr, func(i, j int) bool {
		return (arr[i][0] / arr[i][1]) > (arr[j][0] / arr[j][1])
	})
	var result [][2]float64
	count := 0
	for w != 0 {
		if w >= arr[count][1] {
			result = append(result, arr[count])
			w -= arr[count][1]
		} else if w != 0 {
			result = append(result, [2]float64{arr[count][0] / arr[count][1] * w, w})
			w -= w
		}
		count++
		if count >= len(arr) {
			break
		}
	}
	var sum float64
	for i := range result {
		sum += result[i][0]
	}
	fmt.Printf("%.3f", sum)
}

func ReadItem() (int, float64, [][2]float64) {
	var (
		number int
		weight float64
		p      float64
		n      float64
		result [][2]float64
	)
	fmt.Scan(&number, &weight)
	for i := 0; i < number; i++ {
		fmt.Scan(&p, &n)
		result = append(result, [2]float64{p, n})
	}
	return number, weight, result
}
