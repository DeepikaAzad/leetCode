package easy

import "fmt"

func main() {
	num := []int{5, 5, 5, 10, 20}
	fmt.Println(isChangeable(num)) // true

	num = []int{5, 20}
	fmt.Println(isChangeable(num)) // false

	num = []int{5, 5, 10}
	fmt.Println(isChangeable(num)) // true

	num = []int{10, 10}
	fmt.Println(isChangeable(num)) // false

	num = []int{10, 5}
	fmt.Println(isChangeable(num)) // false
}

func isChangeable(num []int) bool {
	fiveCount, tenCount := 0, 0
	for i := 0; i < len(num); i++ {
		if num[i] == 5 {
			fiveCount++
			// return true
		}

		if num[i] == 10 {
			if fiveCount >= 1 {
				fiveCount--
				tenCount++
				// return true
			} else {
				return false
			}
		}

		if num[i] == 20 {
			if fiveCount > 0 && tenCount > 0 {
				fiveCount--
				tenCount--
				// return true
			} else if fiveCount >= 3 {
				fiveCount -= 3
				// return true
			} else {
				return false
			}

		}
	}

	return true
}