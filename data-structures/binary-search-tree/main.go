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

func main() {
	bst := BinarySearchTree{}
	bst.Insert(100)
	bst.Insert(50)
	bst.Insert(150)
	bst.Insert(300)
	bst.Insert(40)
	fmt.Println(bst.Find(40))
}
