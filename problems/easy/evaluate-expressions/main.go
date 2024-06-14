package main

/*
-1: +
-2: -
-3: /
-4: *

inside of binary search tree, the leafs are the numbers and parents are going to be
the operators and we need the group them together and get to the result
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func main() {

}
//O(n) Time | O(height) Space 
func evaluateExpression(tree BinarySearchTree) int {
	return recursive(tree.root)
}

func recursive(node *Node) int {
	//base case: when got to leaf nodes we do the evaluation
	//since "leafs" are positive numbers and operators are negative
	if node.val > 0 {
		return node.val //just return the number to operator
	}
	//we need result from left and right
	leftValue := recursive(node.left)
	rightValue := recursive(node.right)
	//then we apply the operator on them based on what we have on "node.val"
	if node.val == -1 {
		return leftValue + rightValue
	}
	if node.val == -2 {
		return leftValue - rightValue
	}
	if node.val == -3 {
		return leftValue / rightValue
	}
	if node.val == -4 {
		return leftValue * rightValue
	}
	panic("operation not supported")
}
