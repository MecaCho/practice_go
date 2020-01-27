package algorithms

import (
	"fmt"
	"sort"
)

type Node struct {
	Name   string
	memory int
}

type Nodes []Node

func (n Node) String() string {
	return fmt.Sprintf("Name: %s, memory: %d.", n.Name, n.memory)
}

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].memory < n[j].memory
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func SortNodes() {
	nodes := Nodes{
		{"node1", 100},
		{"node2", 99},
		{"node3", 90},
	}

	sort.Sort(nodes)
	fmt.Println(nodes)
}
