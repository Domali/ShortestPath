package main

import(
    "fmt"
    g "./graph"
)

func main(){
    graph := make(map[string]*g.GraphNode)
    g.ParseFile(graph, "./data")    
    g.Dij(graph)
    g.PrintNodeConn(graph)
    fmt.Println(graph)
}