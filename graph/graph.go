package graph

import(
    s "strings"
    sc "strconv"
    "bufio"
    "fmt"
    "os"
)

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
    node := graph[x]
    nc := graph[y]
    ge := GraphEdge{nc,w}
    node.Connections = append(node.Connections, ge)    
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

func ParseFile(graph map[string]*GraphNode){
    f, err := os.Open("./data")
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

func dij(graph map[string]*GraphNode){
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