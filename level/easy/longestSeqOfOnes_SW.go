package main

import (
	"fmt"
)

func main() {
	num := []int{0, 0, 1, 1, 1, 1, 1, 0, 1, 1}
	fmt.Println(longestOnes(num))
}

func longestOnes(num []int) int {
	start, end := 0, 0
	count, maxLen := 0, 0
	for end < len(num) {
		if num[end] != 1 {
			for count > 0 {
				count--
				start++
			}
			start++
			end++
			continue
		}

		count++

		if maxLen < count {
			maxLen = count
		}
		end++
	}
	return maxLen
}