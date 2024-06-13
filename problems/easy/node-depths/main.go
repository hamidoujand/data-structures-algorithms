package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}
type DepthNode struct {
	node  *Node
	depth int
}

func main() {

}

// O(n) Time | O(height) Space
func nodeDepthsIterative(tree BinaryTree) int {
	sumOfDepths := 0
	stack := []DepthNode{
		{
			node:  tree.root,
			depth: 0,
		},
	}

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sumOfDepths += currentNode.depth
		//check it's left and right exists and if, put them on top of stack
		if currentNode.node.left != nil {

			stack = append(stack, DepthNode{
				node:  currentNode.node.left,
				depth: currentNode.depth + 1,
			})
		}
		if currentNode.node.right != nil {
			stack = append(stack, DepthNode{
				node:  currentNode.node.right,
				depth: currentNode.depth + 1,
			})
		}
	}
	return sumOfDepths
}

// O(n) Time | O(heigh) Space
func nodeDepthsRecursive(tree BinaryTree) int {
	return recursive(tree.root, 0)
}

func recursive(node *Node, currentDepth int) int {
	//handle base case
	if node == nil {
		return 0
	}
	return currentDepth + recursive(node.left, currentDepth+1) + recursive(node.right, currentDepth+1)
}
