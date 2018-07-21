package tree

import (
    "fmt"
)

type Tree struct {
    head *Leaf
}

type leaf struct {
    parent *leaf
    children leaf[]
    data *GraphNode
}

func (t Tree) addLeaf() {
    
}