package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
}


// Improvised
func reverse(x int) int {
	s := 0
	for x != 0 {
		t := x % 10
		x = x / 10
		s = s*10 + t
	}
	return s
}

func reverse2(x int) int {
	numList := make(map[int]int)
	count := 0
	for {
		if x/10 != 0 {
			numList[count] = x % 10
		} else {
			numList[count] = x % 10
			break
		}
		x = x / 10
		count++

	}
	temp := 1
	x = 0
	for i := len(numList) - 1; i >= 0; i-- {
		x = x + numList[i]*temp
		temp = temp * 10
	}

	if x > int(math.Pow(2, 31))-1 || x < int(math.Pow(-2, 31)) {
		return 0
	}
	return x

}
