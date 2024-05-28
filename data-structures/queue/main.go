package main

import "fmt"

/*
Queue: it's an abstract data structure which follows the "FIFO" principles, first in first out.

an example in real world: "printer uses a queue"

Big O:
insertion : O(1)
deletion:  O(1)
*/

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
	size  int
}

func (q *Queue) Enqueue(val int) int {
	node := Node{
		val: val,
	}
	//no item
	if q.first == nil {
		q.first = &node
		q.last = &node
		q.size++
		return q.size
	}

	//otherwise
	oldLast := q.last
	oldLast.next = &node
	q.last = &node
	q.size++
	return q.size
}

func (q *Queue) Dequeue() *Node {
	//no item
	if q.first == nil {
		return nil
	}

	//only 1 item
	if q.first == q.last {
		node := q.first
		q.first = nil
		q.last = nil
		q.size--
		return node
	}

	//otherwise
	oldFirst := q.first
	newFirst := oldFirst.next
	q.first = newFirst
	q.size--
	return oldFirst
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	fmt.Println(q)
	q.Enqueue(2)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
