package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// - move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			// recursively insert
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// - move right
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			// recursively insert
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and Return true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		// - move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// - move left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	// - init tree
	tree := &Node{Key: 100}
	fmt.Println("Tree : ", tree)

	// - insert new key into tree
	tree.Insert(200)
	fmt.Println("Insert some value to the keyTree : ", tree)

	// - search the key
	fmt.Println("Search for 200 : ", tree.Search(200))
}
