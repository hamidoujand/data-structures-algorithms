package main

import (
	"fmt"
	"math"
	"slices"
)

/*
Dijkstra Algorithm: finds the shortest path between vertices
works on weighted graphs

*/

type Vertex struct {
	node   string
	weight int
}

type WeightedGraph struct {
	adjacencyList map[string][]Vertex
}

func (g *WeightedGraph) AddVertex(vertex Vertex) {
	if _, ok := g.adjacencyList[vertex.node]; !ok {
		g.adjacencyList[vertex.node] = []Vertex{}
	}

}

func (g *WeightedGraph) AddEdge(v1, v2 string, weight int) {
	g.adjacencyList[v1] = append(g.adjacencyList[v1], Vertex{node: v2, weight: weight})
	g.adjacencyList[v2] = append(g.adjacencyList[v2], Vertex{node: v1, weight: weight})
}

func (g *WeightedGraph) Dijkstra(start, end string) []string {
	nodes := PriorityQueue{values: make([]Node, 0)}
	distances := make(map[string]float64)
	previous := make(map[string]string)
	var path []string
	//build initial state
	for vertex := range g.adjacencyList {
		if vertex == start {
			distances[vertex] = 0 //for the start the distance to itself is always 0
			//highest priority is "start" vertex
			nodes.Enqueue(vertex, 0)
		} else {
			distances[vertex] = math.Inf(0) //other for right now "+Inf"
			//the priority of others is also "+Inf"
			nodes.Enqueue(vertex, math.Inf(0))
		}
		//first time the "previous" path to each vertex is ""
		previous[vertex] = ""
	}
	//as long as there is something to visit
	for len(nodes.values) > 0 {
		//get the vertex with highest priority
		current := nodes.Dequeue().val

		//we loop it's neighbors and add them into "nodes" priority queue
		for _, neighbor := range g.adjacencyList[current] {
			//take the distance for the vertex and add "its neighbors's weight" to it
			candidate := distances[current] + float64(neighbor.weight)
			//if the new calculated distance is less that whats currently stored for the "neighbor" inside of "distances"
			//we update it to this candidate value
			if candidate < distances[neighbor.node] {
				distances[neighbor.node] = candidate
				//if the candidate is less, we need to mark from which vertex we got to the neighbor
				previous[neighbor.node] = current
				//last we need to enqueue this "neighbor" into "priority queue" with a new priority which is the
				//"candidate" we calculated
				nodes.Enqueue(neighbor.node, candidate)
			}
		}

	}

	// we need to traverse the "previous" map from "end" to "start"
	seeker := end
	for previous[seeker] != "" {
		path = append(path, seeker)
		seeker = previous[seeker]
	}
	path = append(path, start)
	//reverse it
	slices.Reverse(path)
	return path
}

// Most simple and naive priority queue
type PriorityQueue struct {
	values []Node
}

type Node struct {
	val      string
	priority float64
}

func (p *PriorityQueue) Enqueue(val string, priority float64) {
	p.values = append(p.values, Node{val: val, priority: priority})
	//sort it
	slices.SortFunc(p.values, func(a Node, b Node) int {
		return int(a.priority - b.priority)
	})
}

func (p *PriorityQueue) Dequeue() Node {
	val := p.values[0]
	p.values = p.values[1:]
	return val
}

func main() {
	g := WeightedGraph{
		adjacencyList: make(map[string][]Vertex),
	}
	g.AddVertex(Vertex{node: "A"})
	g.AddVertex(Vertex{node: "B"})
	g.AddVertex(Vertex{node: "C"})
	g.AddVertex(Vertex{node: "D"})
	g.AddVertex(Vertex{node: "E"})
	g.AddVertex(Vertex{node: "F"})
	g.AddEdge("A", "B", 4)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "E", 3)
	g.AddEdge("C", "D", 2)
	g.AddEdge("C", "F", 4)
	g.AddEdge("D", "E", 3)
	g.AddEdge("D", "F", 1)
	g.AddEdge("E", "F", 1)

	fmt.Println(g.Dijkstra("A", "E"))
}
