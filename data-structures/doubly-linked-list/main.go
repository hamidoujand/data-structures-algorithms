package main

import "fmt"

/*
Doubly Linked List: it is almost identical to singly linked lists, and we just keep a
reference to the previous node as well. a browser history is a good example of this data structure.
its takes more memory compare to singly linked list for that extra pointer.

Big O:

Insertion: O(1)
Pop/Shift: O(1)
Searching: O(N)
Access: O(N)

*/

type Node struct {
	val  int
	next *Node
	pre  *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *DoublyLinkedList) Push(val int) {
	node := Node{val: val}
	//first item
	if l.head == nil {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	//otherwise

	currentTail := l.tail
	currentTail.next = &node
	node.pre = currentTail
	l.tail = &node //change the tail to be the new node
	l.length++
}

func (l *DoublyLinkedList) Pop() *Node {
	//empty
	if l.head == nil {
		return nil
	}
	//one item
	if l.length == 1 {
		node := l.tail
		//empty it
		l.head = nil
		l.tail = nil
		l.length--
		node.pre = nil
		return node
	}

	//otherwise

	currentTail := l.tail
	newTail := currentTail.pre
	currentTail.pre = nil
	newTail.next = nil
	l.tail = newTail
	l.length--

	return currentTail
}

func (l *DoublyLinkedList) Shift() *Node {
	//empty
	if l.head == nil {
		return nil
	}

	//one item
	if l.length == 1 {
		node := l.head
		l.head = nil
		l.tail = nil
		l.length--
		return node
	}

	//otherwise
	currentHead := l.head
	newHead := currentHead.next
	currentHead.next = nil
	newHead.pre = nil
	l.head = newHead
	l.length--
	return currentHead
}

func (l *DoublyLinkedList) UnShift(val int) {
	node := &Node{val: val}
	//empty
	if l.head == nil {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	//otherwise
	currentHead := l.head
	currentHead.pre = node
	node.next = currentHead
	l.head = node
	l.length++
}

func (l *DoublyLinkedList) Get(idx int) *Node {
	//idx check
	if idx < 0 || idx >= l.length {
		return nil
	}

	mid := l.length / 2
	if idx <= mid {
		//start from head
		current := l.head
		counter := 0
		for counter < idx {
			current = current.next
			counter++
		}

		return current
	} else {
		//start from tail
		current := l.tail
		counter := l.length - 1
		for counter > idx {
			current = current.pre
			counter--
		}
		return current
	}
}

func (l *DoublyLinkedList) Set(idx int, val int) bool {
	node := l.Get(idx)
	if node == nil {
		return false
	}
	node.val = val
	return true
}

func (l *DoublyLinkedList) Insert(idx int, val int) bool {
	//idx check
	if idx < 0 || idx > l.length {
		return false
	}
	//unShift
	if idx == 0 {
		l.UnShift(val)
		return true
	}

	//push
	if idx == l.length {
		l.Push(val)
		return true
	}

	//otherwise
	node := Node{val: val}
	nodeBefore := l.Get(idx - 1)
	node.next = nodeBefore.next
	node.pre = nodeBefore
	nodeBefore.next = &node
	l.length++
	return true
}

func (l *DoublyLinkedList) Remove(idx int) bool {
	// idx check
	if idx < 0 || idx >= l.length {
		return false
	}

	//shift
	if idx == 0 {
		l.Shift()
		return true
	}

	//pop
	if idx == l.length-1 {
		l.Pop()
		return true
	}

	//otherwise
	nodeToRemove := l.Get(idx)
	previous := nodeToRemove.pre
	after := nodeToRemove.next
	previous.next = after
	after.pre = previous
	nodeToRemove.next = nil
	nodeToRemove.pre = nil
	l.length--
	return true
}

func main() {
	list := DoublyLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	list.Push(6)
	list.Push(7)
	list.Push(8)
	list.Insert(3, 4000)
	fmt.Println(list.Get(3))
	list.Remove(4)
	fmt.Println(list.Get(4))
}
