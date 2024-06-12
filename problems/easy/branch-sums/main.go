package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

// O(n) Time | O(n) space  (recursive call stack)
func (b *BinaryTree) branchSums() []int {
	var result []int
	sum := 0
	b.recursive(b.root, sum, &result)
	return result
}

func (b *BinaryTree) recursive(node *Node, sum int, result *[]int) {
	if node == nil {
		return
	}
	//calculate new sum
	sum += node.value
	//add to results when there is no more left or right to this node
	//leaf node
	if node.left == nil && node.right == nil {
		*result = append(*result, sum)
		return
	}

	//otherwise run it on the left and right
	b.recursive(node.left, sum, result)
	b.recursive(node.right, sum, result)
}

func main() {
	tree := BinaryTree{}
	node1 := Node{value: 1}
	node2 := Node{value: 2}
	node3 := Node{value: 3}
	node4 := Node{value: 4}
	node5 := Node{value: 5}
	node6 := Node{value: 6}
	node7 := Node{value: 7}
	//set the root
	tree.root = &node1

	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	node2.right = &node5
	node3.left = &node6
	node3.right = &node7

	results := tree.branchSums()
	fmt.Println(results)
}
