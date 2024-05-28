package main

import "fmt"

/*
Stack: its an abstract data structure which follows "LIFO" last in first out principle.
A great example of this data structure is used as "Call Stack"
another one "undo-redo" functionality

since its an abstraction data structure, there are multiple ways to implement it

Big O:
insertion: O(1)
deletion: O(1)
*/

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	first  *Node
	last   *Node
	length int
}

func (s *Stack) Push(val int) int {
	node := &Node{val: val}
	//first item
	if s.first == nil {
		s.first = node
		s.last = node
		s.length++
		return s.length
	}

	//otherwise
	currentFirst := s.first
	node.next = currentFirst
	s.first = node
	s.length++
	return s.length
}

func (s *Stack) Pop() *Node {
	//no item
	if s.first == nil {
		return nil
	}
	//one item
	if s.first == s.last {
		node := s.first
		s.length--
		s.first = nil
		s.last = nil
		return node
	}

	//otherwise
	currentFirst := s.first
	newFirst := currentFirst.next
	s.first = newFirst
	s.length--
	currentFirst.next = nil
	return currentFirst
}

func main() {
	s := Stack{}
	s.Push(1)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
