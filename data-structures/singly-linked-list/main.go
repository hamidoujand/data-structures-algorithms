package main

import "fmt"

/*
Singly Linked List: it's an ordered list of data just like array or slice but it
does not have any indices and it has "head" "tail" "length" and each node has a
pointer to next element. a node is going to store the value and pointer to next
node.

to access nodes we need to iterate over it from "head" till we find it

they are good at insertion and deletion compare to arrays, because they do not
have an index to update each time we insert or delete.

used when you have a long set of data and you do not need random access, just
insertion or deletion on that set.


Big O:

Insertion: O(1)
Search: O(n)
Access: O(n)
Pop: O(n)
Shift: O(1)
Remove: O(n)
*/

type Node struct {
	val  int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *SinglyLinkedList) Push(val int) {
	node := Node{
		val: val,
	}
	if l.head == nil {
		//its the first value for the "head" and "tail" is going to be this value
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	//otherwise its not the first one, so we find the "tail"
	l.tail.next = &node
	//change the tail
	l.tail = &node
	//increment length
	l.length++
}

func (l *SinglyLinkedList) Pop() *Node {
	//no item
	if l.head == nil {
		return nil
	}
	//only 1 item inside of list
	if l.head == l.tail {
		temp := l.tail
		l.head = nil
		l.tail = nil
		l.length = 0
		return temp
	}
	//more than one item inside of it
	var pre *Node
	current := l.head
	for current.next != nil {
		pre = current
		current = current.next
	}
	//we have the "previous tail"
	l.tail = pre
	pre.next = nil //set its "next" to nil
	l.length--
	//check to see if length is 0
	if l.length == 0 {
		//set the "head" and "tail" to nil
		l.head = nil
		l.tail = nil
	}
	return current
}

func (l *SinglyLinkedList) Shift() *Node {
	// no item
	if l.head == nil {
		return nil
	}

	//one item
	if l.head == l.tail {
		temp := l.head
		l.head = nil
		l.tail = nil
		l.length = 0
		return temp
	}

	//more than one
	temp := l.head
	l.head = temp.next
	l.length--
	if l.length == 0 {
		//set the tail and head to nil
		l.head = nil
		l.tail = nil
	}
	temp.next = nil
	return temp
}

func (l *SinglyLinkedList) UnShift(val int) {
	//empty list
	if l.head == nil {
		l.Push(val)
		return
	}
	node := Node{val: val}
	oldHead := l.head
	//make it new head
	l.head = &node
	//connect the dots
	node.next = oldHead
	l.length++

}

func (l *SinglyLinkedList) Get(idx int) *Node {
	if idx < 0 || idx >= l.length {
		return nil
	}
	//we need to traverse it
	currentIdx := 0
	currentNode := l.head
	for currentIdx < idx {
		currentNode = currentNode.next
		currentIdx++
	}
	//we got the index
	return currentNode
}

func (l *SinglyLinkedList) Insert(idx int, val int) bool {
	if idx < 0 || idx > l.length {
		return false
	}
	//first item
	if idx == 0 {
		l.UnShift(val)
		return true
	}

	//last item
	if idx == l.length {
		l.Push(val)
		return true
	}

	node := Node{val: val}
	//get the item currently at 1 index before
	previous := l.Get(idx - 1)
	node.next = previous.next //attach next item to the new node
	previous.next = &node
	l.length++
	return true
}

func (l *SinglyLinkedList) Set(idx int, val int) bool {
	node := l.Get(idx)
	if node == nil {
		return false
	}
	node.val = val
	return true

}

func (l *SinglyLinkedList) Remove(idx int) bool {
	if idx < 0 || idx >= l.length {
		return false
	}

	//remove first item
	if idx == 0 {
		l.Shift()
		return true
	}
	//remove last item
	if idx == l.length-1 {
		l.Pop()
		return true
	}

	//now in middle
	previousNode := l.Get(idx - 1)
	currentNode := previousNode.next
	nextCandidate := currentNode.next
	//connect the dots
	previousNode.next = nextCandidate
	currentNode.next = nil
	l.length--
	return true
}

func (l *SinglyLinkedList) Reverse() {
	//swap head and tail right away
	head := l.head
	l.head = l.tail
	l.tail = head

	current := head
	var pre *Node
	for i := 0; i < l.length; i++ {
		nextNodeToOperateOn := current.next
		//current is going to now point to whatever "pre" points
		current.next = pre
		pre = current                 //previous now becomes the current node
		current = nextNodeToOperateOn //current moves to be the next candidate to operate on
	}
}

func main() {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Reverse()
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
}
