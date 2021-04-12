package main

import "fmt"

/*
	Here is the implementation of stack in golang
*/

type Stack []int

// IsEmpty will check the value that available in stack or not
func (s *Stack) IsEmpty() bool {
	if len(*s) != 0 {
		return false
	}
	return true
}

// Push will add a value into the first element
func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() bool {
	// - check wether the stack isEmpty
	if s.IsEmpty() == true {
		// - there is no value that need to POP
		return false
	}

	// - remove the last element in that slice
	*s = (*s)[:len(*s)-1]
	return true
}

func (s *Stack) Peek() (top int) {
	return (*s)[len(*s)-1]
}

func main() {
	// init stack
	var stack Stack

	// - push few value into stack
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)

	fmt.Println("Current Stack after push : ", stack)

	// - pop one of the value
	if stack.Pop() == true {
		fmt.Println("Stack after done pop one value : ", stack)
	}

	// printing the top element using peek function
	fmt.Println("Peek Stack : ", stack.Peek())

	if stack.IsEmpty() == true {
		fmt.Println("The Stack is Empty")
	}
}
