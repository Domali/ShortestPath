package main

import (
	"fmt"

	g "github.com/domali/ShortestPath/graph"
)

func main() {
	graph := make(map[string]*g.GraphNode)
	g.ParseFile(graph, "./data2")
	g.Dij(graph)
	g.PrintNodeConn(graph)
	fmt.Println(graph)
}
