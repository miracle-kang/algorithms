package tree

import (
	"container/list"
)

// INode interface
type INode interface {

	// Value of the node
	Value() int

	// Left of the node child
	Left() INode

	// Right of the node child
	Right() INode
}

// LayerPrint a tree
// root: The tree root node
func LayerPrint(root INode) {
	queue := list.New()

	node := root
	queue.PushBack(node)

	for queue.Len() > 0 {
		element := queue.Front()
		node = queue.Remove(element).(INode)
		print(node.Value(), " ")

		if node.Left() != nil {
			queue.PushBack(node.Left())
		}
		if node.Right() != nil {
			queue.PushBack(node.Right())
		}
	}
	println()
}
