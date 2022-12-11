// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"strings"
)

//ccaabbb

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct2("ccaabbb"))
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	strList := strings.Split(s, "")
	m := make(map[string]int)
	for i := 0; i < len(strList); i++ {
		if _, ok := m[strList[i]]; !ok {
			m[strList[i]] = 1
		} else {
			m[strList[i]] += 1
		}

	}
	fmt.Println(m)
	smallest, longest := 0, 0
	for _, val := range m {
		if val > longest {
			smallest = longest
			longest = val
		}
	}

	return smallest + longest

}

func lengthOfLongestSubstringTwoDistinct2(s string) int {
	strList := strings.Split(s, "")
	m := make(map[string]int)
	for i := 0; i < len(strList); i++ {
		if _, ok := m[strList[i]]; !ok {
			m[strList[i]] = 1
		} else {
			m[strList[i]] += 1
		}

	}
	fmt.Println(m)
	keys := make([]int, 0, len(m))
	for _, v := range m {
		keys = append(keys, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Println(keys)
	count, val := 0, 0
	for count < 2 {
		val += keys[count]
		count++
	}
	return val

}
