package main

import (
	"fmt"
)

type LinkNode struct {
	value int
	next *LinkNode
}

func CreateList(a []int) *LinkNode{
	head := new(LinkNode)
	head.value = a[0]
	tmp := head
	for _, i := range a[1:] {
		next := new(LinkNode)
		next.value = i
		tmp.next = next
		tmp = next
	}
	return head
}

func ReverseList(head *LinkNode) *LinkNode{
	if head == nil || head.next == nil{
		return head
	}
	cur := head.next
	pre := head
	head.next = nil
	for cur != nil{
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre

}


func ThroughList(head *LinkNode){
	str := ""
	node := head
	for node != nil{
		str += fmt.Sprintf("->%d", node.value)
		node = node.next
	}
	println(str)
}

func main(){
	node := CreateList([]int{1,2,3,4,5})
	ThroughList(node)
	rev := ReverseList(node)
	fmt.Printf("%#v\n", rev)
	ThroughList(rev)
}
