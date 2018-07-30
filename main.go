package main

import (
	g "github.com/domali/ShortestPath/graph"
)

func main() {
	graph := g.Graph{}
	graph.InitGraph("./data2")
	graph.GenAllSpt()
	graph.PrintNodesSpt("0")
}
