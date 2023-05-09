// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 5, 3, -22, 1000, -10}
	result := maxSubArray(nums, 3)
	fmt.Println(result)
}

func maxSubArray(nums []int, k int) int {
	i, j, max, sum := 0, 0, 0.0, 0
	// var max int64
	for j < len(nums) {
		sum = sum + nums[j]

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			max = math.Max(float64(sum), float64(max))
			fmt.Println(i, j, max, sum)
			sum = sum - nums[i]
			i++
			j++
		}

	}

	for i < len(nums) {
		sum = sum + nums[j-len(nums)]
		max = math.Max(float64(sum), float64(max))
		sum = sum - nums[i]
		j++
		i++

	}

	return int(max)
}
