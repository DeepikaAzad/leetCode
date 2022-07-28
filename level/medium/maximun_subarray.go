// Kadane’s Algorithm

package main

import "fmt"

func main() {
	nums := []int{1, -1}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max_so_far, max_ending_here := nums[0], 0
	for _, val := range nums {
		max_ending_here = max_ending_here + val
		if max_ending_here > max_so_far {
			max_so_far = max_ending_here
		}
		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}

	return max_so_far
}
