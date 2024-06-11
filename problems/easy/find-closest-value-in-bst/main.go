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

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(val float64) {
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
	bst := BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(13)
	bst.Insert(5)
	bst.Insert(22)
	bst.Insert(14)

	fmt.Println(findClosestValueInBSTIterative(bst, 12))
	fmt.Println(findClosestValueInBSTRecursive(bst, 12))
}

// Average: O(log n) Time | O(log n) Space
// Worst: O(n)  Time | O(n) Space    weird linked list style trees
func findClosestValueInBSTRecursive(tree BinarySearchTree, target float64) float64 {
	return traverse(tree.root, target, math.Inf(0))
}

func traverse(node *Node, target float64, candidate float64) float64 {
	if node == nil {
		return candidate
	}

	//if the difference between "target and candidate" is larger than the difference between "target and new node value"
	//we update it since we now found a smaller difference with that new node
	if math.Abs(target-candidate) > math.Abs(target-node.val) {
		candidate = node.val
	}
	//decide which part of tree we moving now
	if target < node.val {
		//left
		return traverse(node.left, target, candidate)
	} else if target > node.val {
		//right
		return traverse(node.right, target, candidate)
	} else {
		return candidate
	}
}

// Average: O(log n) Time | O(1) Space
// Worst: O(n)  Time | O(1) Space
func findClosestValueInBSTIterative(tree BinarySearchTree, target float64) float64 {
	currentNode := tree.root
	candidate := math.Inf(0)

	for currentNode != nil {
		if math.Abs(target-candidate) > math.Abs(target-currentNode.val) {
			candidate = currentNode.val
		}

		//change the currentNode
		if target > currentNode.val {
			//right
			currentNode = currentNode.right
		} else if target < currentNode.val {
			//left
			currentNode = currentNode.left
		} else {
			//equal
			return candidate
		}
	}
	return candidate
}
