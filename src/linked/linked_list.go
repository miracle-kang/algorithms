package linked

// Node is a node for linked list
type Node struct {
	pre  *Node
	next *Node
	data interface{}
}

// List is a list implement by linked
type List struct {
	head *Node
	tail *Node
	size int
}

// will be implement future
