package main

import (
	"fmt"
)

// Node struct will represent the boiler plate of our linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList ...
type LinkedList struct {
	Head   *Node
	Length int
}

// Prepend is a function that similar with the Insert function.
func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.Head
	for l.Length != 0 {
		fmt.Printf("%d ", toPrint.Value)
		toPrint = toPrint.Next
		l.Length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next.Value != value {
		if previousToDelete.Next.Next == nil {
			fmt.Println("Data Not Found")
			return
		}
		previousToDelete = previousToDelete.Next
	}

	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}

func main() {
	myList := LinkedList{}
	node1 := &Node{Value: 70}
	node2 := &Node{Value: 53}
	node3 := &Node{Value: 63}

	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)

	// - data printed before element deleted
	myList.PrintListData()

	// - delete with value 63
	myList.DeleteWithValue(63)
	myList.PrintListData()
}
