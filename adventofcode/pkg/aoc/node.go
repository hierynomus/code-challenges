package aoc

type Node struct {
	Id          string
	Connections []*Node
}

func NewNode(id string) *Node {
	return &Node{
		Id: id,
	}
}

func (n *Node) Connect(node *Node) {
	n.Connections = append(n.Connections, node)
	node.Connections = append(node.Connections, n)
}
