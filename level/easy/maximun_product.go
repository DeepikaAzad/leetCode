
// https://leetcode.com/problems/maximum-product-of-three-numbers/submissions/
package main

import (
	"fmt"
	"sort"
)

func main() {
	max := []int{1, -4, 3, -6, 7, 0}
	fmt.Println(maxiumnProduct(max))
}

func maxiumnProduct(input []int) int {
	// sort the array
	index := len(input) - 1
	sort.Sort(sort.IntSlice(input))
	if input[index]*input[index-1]*input[index-2] > input[0]*input[1]*input[index] {
		return input[index] * input[index-1] * input[index-2]
	}
	return input[0] * input[1] * input[index]
}
