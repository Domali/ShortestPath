package graph

import(
    s "strings"
    sc "strconv"
    "bufio"
    "fmt"
    "os"
)
//Types for shortest path map
type ShortestPath struct {
    SptSet map[string]int
    Nodes []ShortestPathNode
}

type ShortestPathNode struct {
    Id string
    Cost int
    Parent *ShortestPathNode
}

func (s *ShortestPath) Initialize(graph map[string]*GraphNode) {
    s.SptSet = make(map[string]int)
    for node,_ := range graph {
        if node != "start" && node != "end" {

            s.Nodes = append(s.Nodes, ShortestPathNode{Id: node, Cost: -1, Parent: nil})
        }
    }
}

func (s *ShortestPath) PrintSpt() {
    fmt.Println("Shortest Path Table")
    fmt.Println("ID\tCost\tParent")
    for _,node := range s.Nodes {
        var parent string
        if node.Parent == nil {
            parent = "nil"
        } else {
            parent = node.Parent.Id
        }
        fmt.Printf("%v\t%v\t%v\n", node.Id, node.Cost, parent)
    }
}

//Types for making the graph
type GraphNode struct {
    Id string
    Connections []GraphEdge
}

type GraphEdge struct {
    GraphNode *GraphNode
    Weight int
}

func check(e error){
    if e != nil {
        panic(e)
    }
}

func addOneWayEdge (graph map[string]*GraphNode, x string, y string, w int){  
    graph[x].Connections = append(graph[x].Connections, GraphEdge{graph[y],w})
}

func addEdge(graph map[string]*GraphNode, x string, y string, w int){
    addOneWayEdge(graph, x, y, w)
    addOneWayEdge(graph, y, x, w)
}

func setStartEnd(graph map[string]*GraphNode, s string, e string){
    graph["start"] = graph[s]
    graph["end"] = graph[e]
}

func addNode(graph map[string]*GraphNode, id string){
    graph[id] = &GraphNode{Id: id}
}

func ParseFile(graph map[string]*GraphNode, fileName string){
    f, err := os.Open(fileName)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for i := 0; scanner.Scan(); i++ {
        line := scanner.Text()
        if i >= 2 {
            conn := s.Split(line,",")
            w,err := sc.Atoi(conn[2])
            check(err)
            addEdge(graph, conn[0], conn[1], w)
        } else {
            line = line[1:len(line)-1]
            p := s.Split(line,",")
            if i == 1 {
                setStartEnd(graph,p[0],p[1])
            } else if i == 0 {
                for _, id := range p {
                    addNode(graph, id)
                }
            }
        } 
    }
}

func Dij(graph map[string]*GraphNode){
    sp := new(ShortestPath)
    sp.Initialize(graph)
    sp.PrintSpt()
}

func PrintNodeConn(graph map[string]*GraphNode){
    for node,data := range graph {
        if node != "start" && node != "end" {
            fmt.Println()
            fmt.Println(node)
            for _,thing := range data.Connections {
                f := thing.GraphNode
                fmt.Printf("%v,%v,%v\n",node,f.Id, thing.Weight)
            }
        }
    }
}