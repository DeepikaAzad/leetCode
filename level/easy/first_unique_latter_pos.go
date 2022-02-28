package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstUniqChar("aab"))
}

// improvides
func firstUniqChar(s string) int {
	strMap := map[string]int{}
	char := []rune(s)
	for _, v := range char {
		if _, ok := strMap[string(v)]; !ok {
			strMap[string(v)] = 1
		} else {
			strMap[string(v)] = strMap[string(v)] + 1
		}

	}

	for i, v := range char {
		if strMap[string(v)] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	strMap := map[byte]int{}
	for i := range s {
		strMap[s[i]]++
	}

	for i := range s {
		if strMap[s[i]] == 1 {
			return i
		}
	}

	return -1
}

