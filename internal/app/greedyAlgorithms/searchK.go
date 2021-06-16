package greedyalgorithms

import (
	"fmt"
)

func SearchK() {
	var n int
	fmt.Scan(&n)
	if n <= 2 {
		fmt.Println(1)
		fmt.Println(n)
		return
	}
	result := []int{}

	for i := 1; i <= n; i++ {
		if i+1 < n && n-i > i {
			n -= i
			result = append(result, i)
		} else {
			result = append(result, n)
			break
		}
	}
	fmt.Println(len(result))
	for i := range result {
		fmt.Printf("%v ", result[i])
	}
}
