package medium

import (
	"strings"
)

func lengthOfLongestSubstring1(s string) int {
	stringArray := strings.Split(s, "")
	m := make(map[string]int)
	for i := 0; i < len(stringArray); i++ {
		if _, ok := m[stringArray[i]]; !ok {
			m[stringArray[i]] = 1
		} else {
			m[stringArray[i]] += m[stringArray[i]]
		}
	}
	return len(m)
}

func lengthOfLongestSubstring(s string) (l int) {
	stringArray := strings.Split(s, "")
	m := make(map[string]int)
	count := 0
	l = 0
	for i := count; i < len(stringArray); i++ {
		if _, ok := m[stringArray[i]]; !ok {
			m[stringArray[i]] = 1
			if len(m) > l {
				l = len(m)
			}
		} else {
			count = count + 1
			i = count
			m = map[string]int{}
			m[stringArray[count]] = 1
		}
	}
	return l
}
