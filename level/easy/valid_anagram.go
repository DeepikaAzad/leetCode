// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("anagam", "nagaram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	str := map[byte]int{}
	str2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		str[s[i]]++
		str2[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if str[s[i]] != str2[s[i]] {
			return false
		}
	}
	return true
}
