package greedyalgorithms

import (
	"fmt"
	"sort"
)

func SearchPoints() (int, []int) {
	arr := ReadSections()
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	result := []int{}
	count := 0
	for count < len(arr) {
		result = append(result, arr[count][1])
		count++
		for count < len(arr) && (result[len(result)-1] >= arr[count][0] && result[len(result)-1] <= arr[count][1]) {
			count++
		}
	}
	return len(result), result
}

func ReadSections() [][2]int {
	var number int
	var pointOne int
	var pointTwo int

	result := [][2]int{}
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&pointOne, &pointTwo)
		result = append(result, [2]int{pointOne, pointTwo})
	}
	return result
}
