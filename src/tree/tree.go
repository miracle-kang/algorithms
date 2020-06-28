package tree

// Node is a tree node
type Node struct {
	parent *Node
	left   *Node
	right  *Node

	data int
}
