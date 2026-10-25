package main

import "fmt"

func main() {
	var node3 Node = Node{Value: 30, Next: nil}
	var node2 Node = Node{Value: 20, Next: &node3}
	var node1 Node = Node{Value: 10, Next: &node2}
	printList(&node1)
	fmt.Println(sumList(&node1))
	fmt.Println(findNode(&node1, 20))
	printList(nil)
	sumList(nil)
	findNode(nil, 10)
}

type Node struct {
	Value int
	Next  *Node
}

func printList(head *Node) {
	fmt.Println(head)
	if head != nil {
		printList((*head).Next)
	}
}

func sumList(head *Node) int {
	if head == nil {
		return 0
	}
	return (*head).Value + sumList((*head).Next)
}

func findNode(head *Node, value int) *Node {
	if head == nil {
		return nil
	}
	if (*head).Value == value {
		return head
	}
	return findNode(head.Next, value)
}
