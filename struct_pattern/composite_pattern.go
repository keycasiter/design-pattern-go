package struct_pattern

import "fmt"

type Node struct {
	Name       string
	ChildNodes []*Node
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

func (n *Node) AddChildNode(node *Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

func (n *Node) AddChildNodes(nodes []*Node) {
	n.ChildNodes = append(n.ChildNodes, nodes...)
}

func (n *Node) PrintLevel() {
	level := 0
	fmt.Printf("level:%d ,name:%s \n", level, n.Name)
	level++
	for _, childNode := range n.ChildNodes {
		fmt.Printf("level:%d ,name:%s \n", level, childNode.Name)
	}
}
