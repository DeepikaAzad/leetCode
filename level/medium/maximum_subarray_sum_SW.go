// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 5, 1, 8, 2, 9, 1}
	fmt.Println(MaxSumSubArray(arr, 3))
}

func MaxSumSubArray(arr []int, size int) float64 {
	max, windowSum, start := 0.0, 0, 0
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		if (i - start + 1) == size {
			max = math.Max(float64(windowSum), float64(max))
			windowSum -= arr[start]
			start++
		}

	}

	return max
}
