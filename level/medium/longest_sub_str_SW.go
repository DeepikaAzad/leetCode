package main

import (
	"fmt"
)

func main() {
	str := ""

	str = "findlongestsubstring"
	fmt.Println(longestSubstring(str)) // dlongest

	str = "longestsubstr"
	fmt.Println(longestSubstring(str)) // longest

	str = "abbcdafeegh"
	fmt.Println(longestSubstring(str)) // bcdafe

	str = "aaaaaa"
	fmt.Println(longestSubstring(str)) // a
}

func longestSubstring(str string) string {
	// s := []byte(str)
	s := str
	start, end := 0, 0
	maxStart := 0
	maxLen := 0
	count := make(map[string]int)
	for end < len(s) {
		count[string(s[end])]++
		// First match window

		for count[string(s[end])] > 1 {
			count[string(s[start])]--
			start++
		}

		if maxLen < end-start+1 {
			maxLen = end - start + 1
			maxStart = start
		}

		end++
	}
	return s[maxStart : maxStart+maxLen]
}