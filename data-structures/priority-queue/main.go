package main

import "fmt"

/*

Priority Queue: binary heap often used to implement "priority queues", and a priority queue is a data structure
in which each element has a priority associate with it and elements with higher priority are served before elements
with lower priority.

Priority Queue is an ABSTRACT concept, they are separated from HEAPS, you can implement them with array or slice even.



*/

type Node struct {
	val      string
	priority int
}

type PriorityQueue struct {
	values []Node
}

func (p *PriorityQueue) Enqueue(val string, priority int) {
	node := Node{
		val:      val,
		priority: priority,
	}
	//append it to the end
	p.values = append(p.values, node)
	//now bubble it up
	p.bubbleUp()
}

func (p *PriorityQueue) bubbleUp() {
	idx := len(p.values) - 1
	val := p.values[idx]
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		parent := p.values[parentIdx]

		if parent.priority <= val.priority {
			break
		}
		p.values[parentIdx], p.values[idx] = p.values[idx], p.values[parentIdx]
		idx = parentIdx
	}
}

func (p *PriorityQueue) Dequeue() Node {

	p.values[0], p.values[len(p.values)-1] = p.values[len(p.values)-1], p.values[0]
	min := p.values[len(p.values)-1]
	p.values = p.values[:len(p.values)-1]
	if len(p.values) > 1 {
		p.sink()
	}

	return min
}

func (p *PriorityQueue) sink() {
	idx := 0
	length := len(p.values)
	element := p.values[0]

	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2

		var leftChild Node
		var rightChild Node
		var swapIdx *int //nil by default
		if leftChildIdx < length {
			leftChild = p.values[leftChildIdx]
			if leftChild.priority < element.priority {
				swapIdx = &leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = p.values[rightChildIdx]

			if (swapIdx == nil && rightChild.priority < element.priority) || (swapIdx != nil && rightChild.priority < leftChild.priority) {
				swapIdx = &rightChildIdx
			}
		}

		if swapIdx == nil {
			break
		}
		p.values[idx], p.values[*swapIdx] = p.values[*swapIdx], p.values[idx]
		idx = *swapIdx
	}
}

func main() {
	p := PriorityQueue{
		values: make([]Node, 0),
	}
	p.Enqueue("common cold", 5)
	p.Enqueue("gunshot wound", 1)
	p.Enqueue("fever", 4)
	p.Enqueue("broken arm", 2)
	p.Enqueue("stomach ache", 3)
	fmt.Println(p)
	fmt.Println(p.Dequeue())
	fmt.Println(p.Dequeue())
	fmt.Println(p.Dequeue())
}
