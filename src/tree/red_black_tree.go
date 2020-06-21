package tree

// Node is a tree node
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	data   interface{}
}

// Tree is an red-black tree implemention
type Tree struct {
	root *Node
	size int
}
