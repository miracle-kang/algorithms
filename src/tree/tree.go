package tree

import "container/list"

// Node is a tree node
type Node struct {
	parent *Node
	left   *Node
	right  *Node

	data int
}

// Value of the node
func (n *Node) Value() int {
	return n.data
}

// LayerPrint a tree
// root: The tree root node
func LayerPrint(root *Node) {
	queue := list.New()

	node := root
	queue.PushBack(node)

	for queue.Len() > 0 {
		element := queue.Front()
		node = queue.Remove(element).(*Node)
		print(node.data, ",")

		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
}
