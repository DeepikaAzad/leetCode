package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/../"))
}

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	ele := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ele, true

}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func simplifyPath(path string) string {
	s := &stack{}
	dirs := strings.Split(path, "/")
	for _, dir := range dirs {
		if dir == "." || dir == "" || (dir == ".." && s.isEmpty()) {
			continue
		}
		if dir == ".." {
			s.pop()
			continue
		}

		s.push(dir)
	}

	return "/" + strings.Join((*s), "/")
}
