package fibonacci

import (
	"fmt"
)

// Search the last digit in number
func SearchLast() {

	var n uint64
	fmt.Scan(&n)
	if n%100 == 0 {
		fmt.Print(n)
		return
	}
	fmt.Print(n % 10)
}

// Дано число n
// необходимо найти последнюю цифру n-го числа Фибоначчи.
func CalcResidue() {
	var n int
	fibArr := []int{0, 1}
	fmt.Scan(&n)

	if n == 1 {
		fmt.Print(0)
		return
	} else if n == 2 {
		fmt.Print(1)
		return
	}

	for i := 2; i <= n; i++ {
		fibArr = append(fibArr, (fibArr[i-1]+fibArr[i-2])%10)
	}
	fmt.Print(fibArr[len(fibArr)-1])
}


