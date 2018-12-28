package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) print() {
	out := ""
	for n != nil {
		out += fmt.Sprintf(`%d-> `, n.Val)
		n = n.Next
	}
	fmt.Println(out)
}

func reverse(list *Node) *Node {
	if list.Next == nil || list == nil {
		return list
	}
	pre, cur := list, list.Next
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

func main() {
	a := &Node{1, &Node{2, &Node{3, nil}}}
	a.print()
	b := reverse(a)
	b.print()
}
