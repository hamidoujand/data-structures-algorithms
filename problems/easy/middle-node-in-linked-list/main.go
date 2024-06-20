package main

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

}

// O(n) Time  | O(1) Space
func middleNodeInLinkedList(list *LinkedList) *Node {
	count := 0

	current := list.head
	for current != nil {
		count++
		current = current.next
	}

	//find the middle index
	idx := count / 2

	//loop till middle
	middleNode := list.head
	for i := 0; i < idx; i++ {
		middleNode = middleNode.next
	}
	return middleNode
}

// O(n) Time  | O(1) Space
func middleNodeInLinkedListV2(list *LinkedList) *Node {
	//uses 2 pointers pattern
	//if we increase a pointer by 1 and another by 2 twice of the first one
	//then when the second pointers reaches end or nil value that means the first
	//pointer will be in the middle
	slowPointer := list.head
	fastPointer := list.head

	for fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
	}

	return slowPointer
}
