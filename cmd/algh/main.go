package main

import (
	"fmt"

	greedyalgorithms "github.com/NetworkPy/algh/internal/app/greedyAlgorithms"
)

func main() {

	var TreeHuffman greedyalgorithms.TreeHuffman
	TreeHuffman.Map = greedyalgorithms.Freq("aaaaaaaaabcbbbsbxbsbxcbi")

	for _, val := range []int{2, 5, 6, 19, 28, 19, 5, 1, 228, 225} {
		TreeHuffman.Insert([]int{val, 0})
	}
	fmt.Println(TreeHuffman.Tree)
	for i := 0; i < 10; i++ {
		fmt.Println(TreeHuffman.GetMin())
	}

	fmt.Println(TreeHuffman.Tree)
}
