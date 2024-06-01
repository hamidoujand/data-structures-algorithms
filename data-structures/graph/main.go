package main

import "fmt"

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


Graph Traversal:

-DFS: means we first visit the neighbors of a node and neighbors of that node and so on
-BSF: means we visit all neighbors in the same depth first before moving to neighbors of neighbors

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

func (g *Graph) DFSRecursive(start string) []string {
	results := make([]string, 0)
	visited := make(map[string]bool) //can use "empty struct" as well
	g.traverse(start, &results, visited)
	return results
}

func (g *Graph) traverse(vertex string, result *[]string, visited map[string]bool) {
	if vertex == "" {
		return
	}

	visited[vertex] = true
	*result = append(*result, vertex)
	neighbors := g.adjacencyList[vertex]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			g.traverse(neighbor, result, visited)
		}
	}
}

func (g *Graph) DFSIterative(start string) []string {
	results := make([]string, 0)
	visited := make(map[string]bool)
	stack := []string{start} //push start on top of the stack

	visited[start] = true //mark it as visited

	for len(stack) > 0 {
		nextVertex := stack[len(stack)-1] //access
		stack = stack[:len(stack)-1]      //pop it

		//append it into results
		results = append(results, nextVertex)

		neighbors := g.adjacencyList[nextVertex]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}

	}
	return results
}

func (g *Graph) BFSIterative(start string) []string {
	results := make([]string, 0)
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		results = append(results, current)
		visited[current] = true
		for _, neighbor := range g.adjacencyList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return results
}

func main() {
	g := Graph{
		adjacencyList: make(map[string][]string),
	}
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "F")

	fmt.Println(g.DFSRecursive("A"))
	fmt.Println(g.DFSIterative("A"))
	fmt.Println(g.BFSIterative("A"))
}
