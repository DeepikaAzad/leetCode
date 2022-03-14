// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	str := "4193 with words"
	fmt.Println(stringToInt(str))
}

func stringToInt(str string) int {
	str = strings.TrimSpace(str)
	n := len(str)
	// If string is empty return 0
	if n == 0 {
		return 0
	}
	signMultiplier := 1
	start := 0
	if str[0] == '-' {
		signMultiplier = -1
		start = 1
	} else if str[0] == '+' {
		start = 1
	}
	res := 0

	for i := start; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			return res
		}

		res = res*10 + int(str[i]-'0')
		// Check if result exceeds MinInt32 or MaxInt32
		if res*signMultiplier <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultiplier >= math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * signMultiplier
}
