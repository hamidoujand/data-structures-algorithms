package main

import "fmt"

/*
Tree: its a data structure that consists of nodes in a "parent/child" relationship
trees are non-linear data structures they can branch

rules for trees:
- a node can only point to a child, not a sibling or a parent, sibling is a group of nodes with same parent.
- only 1 root is allowed, root is a top node in a tree.

leaf node: a node with no children
edge: connection between each node and parent is called edge.


real examples of trees:
- HTML DOM
- Network Routing
- Abstract Syntax Tree
- AI
- Folders in OS


Kinds of Trees:

Binary Tree: each node can have a most 2 children

Binary Search Tree: they are special case of "binary trees" which data are sorted in particular way.
all values less that root are located to the "left" side and all values greater than "root" are going to be on "right" side
this rule applies to all parent nodes as well.

Big O:

insertion: O(log n)
searching: O(log n)


Tree Traversal: is a way that we are going to visit each node at least 1 time.
this can be apply of any kind of tree, binary ,binary search ...

there are 2 common ways of traversing a tree:

-Breadth first search: traversing across the tree horizontally, means visiting all nodes
in the same level first then move down one layer.

-depth first search: traversing down the tree vertically till end then come up and visit others


when to use "DFS" vs "BSF":

their "time complexity" is the same, so they differ in "space complexity"
and that depends on the "tree" itself, for example if you have a "wide" tree
using "bfs" on it requires a queue that is going to store a lot of data and grows
in this case using "dfs" is better since the depth of the tree is not much compare to it's breadth.

and if you have a tree which is "deep" and not that "wide", using "bfs" on it will take less space


*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(val int) {
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

func (bst *BinarySearchTree) Find(val int) *Node {
	//no root
	if bst.root == nil {
		return nil
	}

	current := bst.root
	for {
		//check to see if current is the target
		if current.val == val {
			return current
		}

		//otherwise we move either to left or right
		if val > current.val {
			//move to right if there is a right
			if current.right == nil {
				return nil
			} else {
				current = current.right
			}
		} else {
			//move to the left if there is a left
			if current.left == nil {
				return nil
			} else {
				current = current.left
			}
		}
	}
}

// Iterative
func (bst *BinarySearchTree) BFS() []int {
	var results []int
	var queue []*Node //queue to hold the visited nodes
	//we add "root" into the queue
	queue = append(queue, bst.root)

	//now we loop till nothing inside of the queue
	for len(queue) > 0 {
		//we access the first from the queue
		candidate := queue[0]
		//put the value into result slice
		results = append(results, candidate.val)
		//check if there is a left or right
		if candidate.left != nil {
			queue = append(queue, candidate.left)
		}
		if candidate.right != nil {
			queue = append(queue, candidate.right)
		}
		//now we dequeue the first from the queue
		queue = queue[1:]
	}

	return results
}

// Recursive
func (bst *BinarySearchTree) DFSPreOrder() []int {
	var result []int
	current := bst.root
	traversePreOrder(current, &result)
	return result
}

func (bst *BinarySearchTree) DFSPostOrder() []int {
	var result []int
	current := bst.root
	traversePostOrder(current, &result)
	return result
}

func (bst *BinarySearchTree) DFSInOrder() []int {
	var result []int
	current := bst.root
	traverseInOrder(current, &result)
	return result
}

func traversePreOrder(node *Node, result *[]int) {
	//end of the branch
	if node == nil {
		return
	}

	//add it into result
	*result = append(*result, node.val)

	if node.left != nil {
		//call it on the left side
		traversePreOrder(node.left, result)
	}
	if node.right != nil {
		traversePreOrder(node.right, result)
	}
}

func traversePostOrder(node *Node, result *[]int) {
	//end of the branch
	if node == nil {
		return
	}

	if node.left != nil {
		//call it on the left side
		traversePostOrder(node.left, result)
	}
	if node.right != nil {
		traversePostOrder(node.right, result)
	}

	//add it into result after exploring it's left and right
	*result = append(*result, node.val)
}

func traverseInOrder(node *Node, result *[]int) {
	//end of the branch
	if node == nil {
		return
	}

	if node.left != nil {
		//call it on the left side
		traverseInOrder(node.left, result)
	}

	//add it into result after

	*result = append(*result, node.val)

	if node.right != nil {
		traverseInOrder(node.right, result)
	}
}

func main() {
	bst := BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(8)
	bst.Insert(20)
	fmt.Println(bst.BFS())
	fmt.Println(bst.DFSPreOrder())
	fmt.Println(bst.DFSPostOrder())
	fmt.Println(bst.DFSInOrder())
}
