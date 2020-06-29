package tree

const (
	// RED color for rb-node
	RED = true
	// BLACK color for rb-node
	BLACK = false
)

// RBNode is red-black tree rb-node
type RBNode struct {
	parent, left, right *RBNode
	color               bool
	data                int
}

// Value of the node
func (rbn *RBNode) Value() int {
	return rbn.data
}

// Left of the node child
func (rbn *RBNode) Left() INode {
	return rbn.left
}

// Right of the node child
func (rbn *RBNode) Right() INode {
	return rbn.right
}

func parent(rbn *RBNode) *RBNode {
	if rbn == nil {
		return nil
	}
	return rbn.parent
}

func grandParent(rbn *RBNode) *RBNode {
	return parent(parent(rbn))
}

func sibling(rbn *RBNode) *RBNode {
	p := parent(rbn)
	if p == nil {
		return nil
	}

	if p.left == rbn {
		return p.right
	}
	return p.left
}

func uncle(rbn *RBNode) *RBNode {
	return sibling(parent(rbn))
}

// RotateLeft by specified node
func rotateLeft(node *RBNode) {
	newNode := node.right
	p := parent(node)
	if newNode == nil {
		return
	}

	node.right = newNode.left
	newNode.left = node
	node.parent = newNode
	if node.right != nil {
		node.right.parent = node
	}

	if p != nil {
		if node == p.left {
			p.left = newNode
		} else {
			p.right = newNode
		}
	}
	newNode.parent = p
}

// RotateRight by specified node
func rotateRight(node *RBNode) {
	newNode := node.left
	p := parent(node)
	if newNode == nil {
		return
	}

	node.left = newNode.right
	newNode.right = node
	node.parent = newNode
	if node.left != nil {
		node.left.parent = node
	}

	if p != nil {
		if node == p.left {
			p.left = newNode
		} else {
			p.right = newNode
		}
	}
	newNode.parent = p
}

// RBTree is a red-black tree implemention
type RBTree struct {
	root *RBNode
	size int
}

// NewRBTree initialize a red-black tree
func NewRBTree() *RBTree {
	return &RBTree{}
}

// Find a value from red-black tree
func (rbt *RBTree) Find(value int) *RBNode {
	node := rbt.root
	for node != nil {
		if value < node.data {
			node = node.left
		} else if value > node.data {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

// Insert a value to red-black tree
func (rbt *RBTree) Insert(value int) {
	newNode := &RBNode{color: RED, data: value}
	node := rbt.root

	// Insert
	for node != nil {
		if value < node.data {
			if node.left == nil {
				newNode.parent = node
				node.left = newNode
			}
			node = node.left
		} else if value > node.data {
			if node.right == nil {
				newNode.parent = node
				node.right = newNode
			}
			node = node.right
		} else {
			break
		}
	}

	// Repair the tree
	insertRepairTree(node)

	// Find the new root
	root := node
	for root.parent != nil {
		root = root.parent
	}
	rbt.root = root
}

// Repair the tree, return the new root
func insertRepairTree(node *RBNode) {

}

func insertCase1(node *RBNode) {

}

func insertCase2(node *RBNode) {

}

func insertCase3(node *RBNode) {

}

// Delete a value from red-black tree
func (rbt *RBTree) Delete(value int) {

}
