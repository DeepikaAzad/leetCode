// SW -> sliding window
package main

import (
	"fmt"
)

func main() {
	str, k := "", 0

	str, k = "abcbdbdbbdcdabd", 2
	fmt.Println(longestSubstring(str, k)) // bdbdbbd

	str, k = "abcbdbdbbdcdabd", 3
	fmt.Println(longestSubstring(str, k)) // bcbdbdbbdcd

	str, k = "abcbdbdbbdcdabd", 5
	fmt.Println(longestSubstring(str, k)) // ""

	str, k = "abcbdbdbbdcdabd", 4
	fmt.Println(longestSubstring(str, k)) // abcbdbdbbdcdabd
}

func longestSubstringKthDistictChar(str string, k int) string {
	// s := []byte(str)
	s := str
	start, end := 0, 0
	maxStart := 0
	maxLen := 0
	count := make(map[string]int)
	for end < len(s) {
		count[string(s[end])]++
		// First match window
		if len(count) == k {
			if maxLen < end-start+1 {
				maxLen = end - start + 1
				maxStart = start
			}
		}
		// maintain window size
		for len(count) > k {
			// fmt.Println(string(s[start]))
			count[string(s[start])]--
			if count[string(s[start])] == 0 {
				delete(count, string(s[start]))
			}
			start++
		}
		end++
	}
	return s[maxStart : maxStart+maxLen]
}