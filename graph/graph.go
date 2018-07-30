package graph

import (
	"bufio"
	"fmt"
	"os"
	sc "strconv"
	s "strings"
)

//Types for dijkstra's shortest path algo
type shortestPath struct {
	sptSet map[string]int
	nodes  map[string]*shortestPathNode
}

type shortestPathNode struct {
	id     string
	cost   int
	parent *shortestPathNode
}

func (s *shortestPath) initialize(graph map[string]*GraphNode) {
	s.sptSet = make(map[string]int)
	s.nodes = make(map[string]*shortestPathNode)
	for node := range graph {
		if node != "start" && node != "end" {
			s.nodes[node] = &shortestPathNode{id: node, cost: -1, parent: nil}
		}
	}
}

func (s *shortestPath) printSpt() {
	fmt.Println("Shortest Path Table")
	fmt.Println("ID\tCost\tParent")
	for _, node := range s.nodes {
		var parent string
		if node.parent == nil {
			parent = "nil"
		} else {
			parent = node.parent.id
		}
		fmt.Printf("%v\t%v\t%v\n", node.id, node.cost, parent)
	}
}

func (s *shortestPath) updateNodeSpt(graph map[string]*GraphNode, curNode string) {
	s.sptSet[curNode] = -1
	curWeight := s.nodes[curNode].cost
	if curWeight == -1 {
		curWeight = 0
	}
	for _, edge := range graph[curNode].connections {
		edgeNode := edge.GraphNode
		if _, ok := s.sptSet[edgeNode.id]; !ok {
			pathCost := curWeight + edge.weight
			if cost := s.nodes[edgeNode.id].cost; cost == -1 || cost > pathCost {
				s.nodes[edgeNode.id].cost = pathCost
				s.nodes[edgeNode.id].parent = s.nodes[curNode]
			}
		}
	}
}

func (s *shortestPath) getNextNode() string {
	var nextNode string
	lowestCost := -1
	for _, node := range s.nodes {
		if _, ok := s.sptSet[node.id]; !ok {
			if (lowestCost == -1 || node.cost < lowestCost) && node.cost != -1 {
				lowestCost = node.cost
				nextNode = node.id
			}
		}
	}
	return nextNode
}

func (s *shortestPath) generateSpt(graph map[string]*GraphNode) {
	nextNode := graph["start"].id
	for {
		s.updateNodeSpt(graph, nextNode)
		if nextNode = s.getNextNode(); nextNode == "" {
			break
		}
	}
}

func Dij(graph map[string]*GraphNode) {
	sp := new(shortestPath)
	sp.initialize(graph)
	sp.generateSpt(graph)
	sp.printSpt()
}

//Types for making the graph
type Graph struct {
	g map[string]*GraphNode
}

type GraphNode struct {
	id          string
	connections []graphEdge
}

type graphEdge struct {
	GraphNode *GraphNode
	weight    int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addOneWayEdge(graph map[string]*GraphNode, x, y string, w int) {
	graph[x].connections = append(graph[x].connections, graphEdge{graph[y], w})
}

func addEdge(graph map[string]*GraphNode, x, y string, w int) {
	addOneWayEdge(graph, x, y, w)
	addOneWayEdge(graph, y, x, w)
}

func setStartEnd(graph map[string]*GraphNode, s, e string) {
	graph["start"] = graph[s]
	graph["end"] = graph[e]
}

func addNode(graph map[string]*GraphNode, id string) {
	graph[id] = &GraphNode{id: id}
}

func ParseFile(graph map[string]*GraphNode, fileName string) {
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i >= 2 {
			conn := s.Split(line, ",")
			w, err := sc.Atoi(conn[2])
			check(err)
			addEdge(graph, conn[0], conn[1], w)
		} else {
			line = line[1 : len(line)-1]
			p := s.Split(line, ",")
			if i == 1 {
				setStartEnd(graph, p[0], p[1])
			} else if i == 0 {
				for _, id := range p {
					addNode(graph, id)
				}
			}
		}
	}
}

func PrintNodeConn(graph map[string]*GraphNode) {
	for node, data := range graph {
		if node != "start" && node != "end" {
			fmt.Println()
			fmt.Println(node)
			for _, thing := range data.connections {
				f := thing.GraphNode
				fmt.Printf("%v,%v,%v\n", node, f.id, thing.weight)
			}
		}
	}
}
