package graph

import (
	"bufio"
	"fmt"
	"os"
	sc "strconv"
	s "strings"
)

type Graph struct {
	graph   map[string]*node
	nodeSpt map[string]*spTree
}

type node struct {
	id          string
	connections []edge
}

type edge struct {
	GraphNode *node
	weight    int
}

type spTree struct {
	sptSet map[string]int
	nodes  map[string]*sptNode
}

type sptNode struct {
	id     string
	cost   int
	parent *sptNode
}

func (g *Graph) initTree(head string) {
	g.nodeSpt[head] = &spTree{}
	g.nodeSpt[head].sptSet = make(map[string]int)
	g.nodeSpt[head].nodes = make(map[string]*sptNode)
	for node := range g.graph {
		g.nodeSpt[head].nodes[node] = &sptNode{id: node, cost: -1, parent: nil}
	}
}
func (g *Graph) PrintNodesSpt(node string) {
	g.nodeSpt[node].printSpt()
}

func (s *spTree) printSpt() {
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

func (g *Graph) updateNodeSpt(curNode, head string) {
	g.nodeSpt[head].sptSet[curNode] = -1
	curWeight := g.nodeSpt[head].nodes[curNode].cost
	if curWeight == -1 {
		curWeight = 0
	}
	for _, edge := range g.graph[curNode].connections {
		edgeNode := edge.GraphNode
		if _, ok := g.nodeSpt[head].sptSet[edgeNode.id]; !ok {
			pathCost := curWeight + edge.weight
			if cost := g.nodeSpt[head].nodes[edgeNode.id].cost; cost == -1 || cost > pathCost {
				g.nodeSpt[head].nodes[edgeNode.id].cost = pathCost
				g.nodeSpt[head].nodes[edgeNode.id].parent = g.nodeSpt[head].nodes[curNode]
			}
		}
	}
}

func (g *Graph) getNextNode(head string) string {
	var nextNode string
	lowestCost := -1
	for _, node := range g.nodeSpt[head].nodes {
		if _, ok := g.nodeSpt[head].sptSet[node.id]; !ok {
			if (lowestCost == -1 || node.cost < lowestCost) && node.cost != -1 {
				lowestCost = node.cost
				nextNode = node.id
			}
		}
	}
	return nextNode
}

func (g *Graph) generateSpt(n string) {
	head := n
	nextNode := n
	for {
		g.updateNodeSpt(nextNode, head)
		if nextNode = g.getNextNode(head); nextNode == "" {
			break
		}
	}
}

func (g *Graph) GenAllSpt() {
	for node := range g.graph {
		g.genNodeSpt(node)
	}
}
func (g *Graph) genNodeSpt(node string) {
	g.initTree(node)
	g.generateSpt(node)
}

func (g *Graph) addOneWayEdge(x, y string, w int) {
	g.graph[x].connections = append(g.graph[x].connections, edge{g.graph[y], w})
}

func (g *Graph) addEdge(x, y string, w int) {
	g.addOneWayEdge(x, y, w)
	g.addOneWayEdge(y, x, w)
}

func (g *Graph) addNode(id string) {
	g.graph[id] = &node{id: id}
}

func (g *Graph) InitGraph(fileName string) {
	g.graph = make(map[string]*node)
	g.nodeSpt = make(map[string]*spTree)
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
			g.addEdge(conn[0], conn[1], w)
		} else {
			line = line[1 : len(line)-1]
			p := s.Split(line, ",")
			if i == 1 {
			} else if i == 0 {
				for _, id := range p {
					g.addNode(id)
				}
			}
		}
	}
}

func (g *Graph) PrintNodeConn() {
	for node, data := range g.graph {
		fmt.Println()
		fmt.Println(node)
		for _, thing := range data.connections {
			f := thing.GraphNode
			fmt.Printf("%v,%v,%v\n", node, f.id, thing.weight)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
