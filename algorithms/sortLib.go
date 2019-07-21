package algorithms

import (
	"fmt"
	"sort"
)

type Node struct {
	Name string
	memory int
}

func (n Node) String() string{
	return fmt.Sprintf("Name: %s, memory: %d.", n.Name, n.memory)
}

func SortNodes()  {

	type NodeList []Node

	nodes := []Node{
		{"node1", 100},
		{"node2", 99},
		{"node3", 90},
	}

	sort.Sort(nodes)

}

