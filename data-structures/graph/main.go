package main

import (
	"fmt"
)

/*
Graphs: it's a collection of "nodes" and "connection" between those nodes

Terminology:

-Vertex: a node
-Edge: connection between nodes

Types Of Graphs:
-Directed: there is a direction between vertices
-Undirected: there is no direction between vertices, its a 2 ways connection

-Weighted: there is a value associated with each edge.
-Unweighted: there is no value associated with each edge.

they have a lot of usage in real world, like "GPS" or "Social Networks" or "Recommendation engines" ...

NOTE: a tree is a type of "Graph"

*/

// Graph is just a simple no-error handling representation of a graph
type Graph struct {
	adjacencyList map[string][]string
}

func (g *Graph) AddVertex(ver string) {
	g.adjacencyList[ver] = make([]string, 0)
}

func (g *Graph) AddEdge(v1, v2 string) {
	g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
}

func (g *Graph) RemoveEdge(v1, v2 string) {
	var v1Filtered []string
	var v2Filtered []string

	for _, vertex := range g.adjacencyList[v1] {
		if vertex != v2 {
			v1Filtered = append(v1Filtered, vertex)
		}
	}

	for _, vertex := range g.adjacencyList[v2] {
		if vertex != v1 {
			v2Filtered = append(v2Filtered, vertex)
		}
	}
	g.adjacencyList[v1] = v1Filtered
	g.adjacencyList[v2] = v2Filtered
}

func (g *Graph) RemoveVertex(v string) {
	for _, vertex := range g.adjacencyList[v] {
		//we remove the edge between those and the 'v'
		g.RemoveEdge(v, vertex)
	}

	//at the end delete from map
	delete(g.adjacencyList, v)
}

func main() {
	g := Graph{
		adjacencyList: make(map[string][]string),
	}
	g.AddVertex("tokyo")
	g.AddVertex("dallas")
	g.AddVertex("aspen")
	g.AddEdge("dallas", "tokyo")
	g.AddEdge("aspen", "dallas")
	g.AddVertex("hongkong")
	g.AddEdge("hongkong", "tokyo")
	g.AddEdge("hongkong", "aspen")
	g.AddEdge("hongkong", "dallas")
	fmt.Println(g)
	g.RemoveVertex("hongkong")
	fmt.Println(g)
}
