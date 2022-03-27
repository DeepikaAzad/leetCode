// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(reverse("hello"))
}

func reverse(str string) string {
	s := []byte(str)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)
}
