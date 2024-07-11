package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   float64
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(val float64) {
	node := Node{val: val}
	//no item in tree
	if bst.root == nil {
		bst.root = &node
		return
	}

	//otherwise
	//start from root
	current := bst.root
	for {
		//for duplicates in any tree is different here we ignore them. you can add a "frequency" into the "NODE"
		//and increase that when a duplicated value was inserted.
		if node.val == current.val {
			return
		}
		//right
		if node.val > current.val {
			//check if something is in the right side
			if current.right == nil {
				//insert it
				current.right = &node
				return
			} else {
				//if there is something is right, change current to be that node
				current = current.right
			}

		} else if node.val < current.val { //left
			//check if there is no node to the left
			if current.left == nil {
				//insert it
				current.left = &node
				return
			} else {
				//change the current to be the node in "left"
				current = current.left
			}
		}
	}

}

func main() {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(14)

	fmt.Println(validateBST(&bst))
}

// O(n) Time | O(d) Space: "d" depth of tree on call stack
func validateBST(tree *BST) bool {
	return recursion(tree.root, math.Inf(-1), math.Inf(1))
}

func recursion(node *Node, min, max float64) bool {
	//leaf
	if node == nil {
		return true
	}

	//node can not be less than min or greater or equal to the max
	if node.val < min || node.val >= max {
		return false
	}

	//current node is valid, now let's check it's left and right
	isLeftValid := recursion(node.left, min, node.val)   // left side can not be greater or equal to the current
	isRightValid := recursion(node.right, node.val, max) //right side must be greater or equal to the current

	return isLeftValid && isRightValid
}
