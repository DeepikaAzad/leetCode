// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	tree := &BinaryTree{}
	tree.insert(100)
	tree.insert(50)
	tree.insert(60)
	tree.insert(110)
	tree.root.traverse(0, "M")
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (b *BinaryTree) insert(data int) {
	if b.root == nil {
		root := &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
		b.root = root
		fmt.Println("[BinaryTree][insert]:", b.root.data)
	} else {
		b.root.insert(data)
	}
}

func (n *BinaryNode) insert(data int) {
	if (n == &BinaryNode{}) {
		return
	}
	if n.data > data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
			fmt.Println("[BinaryNode][insert]:", n.left.data)
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
			fmt.Println("[BinaryNode][insert]:", n.right.data)
		} else {
			n.right.insert(data)
		}
	}
}

func (b *BinaryNode) traverse(n int, str string) {
	if b == nil {
		return
	}
	for i := 0; i < n; i++ {
		//fmt.Fprint(w, " ")
		fmt.Print(" ")
	}

	//fmt.Fprintln(w, "%s:%s", str, b.data)
	fmt.Println("%s:%s", str, b.data)
	b.left.traverse(n+2, "L")
	b.right.traverse(n+2, "R")

}
