package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Push(val int) {
	node := Node{
		val: val,
	}
	if l.head == nil {
		//its the first value for the "head" and "tail" is going to be this value
		l.head = &node

		return
	}
	current := l.head
	var previous *Node
	for current != nil {
		previous = current
		current = current.next
	}
	previous.next = &node
}

func (l *LinkedList) Log() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ---> ", current.val)
		current = current.next
		if current == nil {
			fmt.Printf("nil\n")
		}
	}
}

func main() {
	l := LinkedList{}
	l.Push(1)
	l.Push(1)
	l.Push(3)
	l.Push(4)
	l.Push(4)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Push(6)
	l.Log()
	removeDuplicatesFromLinkedList(&l)
	l.Log()
}

// sorted linked list is required
// O(n) Time | O(1) Space
func removeDuplicatesFromLinkedList(list *LinkedList) {
	currentNode := list.head
	for currentNode != nil {
		nextNode := currentNode.next
		//now we need to find the next distinct node by skipping all duplicates
		for nextNode != nil && nextNode.val == currentNode.val { //this "Loop" is not going to iterate over the same values twice so its not O(n^2)
			//move the next node over to the distinct one
			nextNode = nextNode.next
		}
		//after the loop means we found a node which is distinct
		//we need to delete whatever comes between "current" --> "distinct"
		currentNode.next = nextNode
		currentNode = nextNode
	}
}
