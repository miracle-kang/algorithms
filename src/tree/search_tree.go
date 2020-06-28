package tree

// SearchTree a binary search tree implemention
type SearchTree struct {
	root *Node
	size int
}

// NewSearchTree new and initialize a binary search tree
func NewSearchTree() *SearchTree {
	return &SearchTree{}
}

// Find a value from binary search tree
func (st *SearchTree) Find(value int) *Node {

	node := st.root
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

// Insert a value to binary search tree
func (st *SearchTree) Insert(value int) {
	newNode := &Node{data: value}
	if st.root == nil {
		st.root = newNode
		return
	}

	node := st.root
	for node != nil {
		if value < node.data {
			// insert if left node is null
			if node.left == nil {
				node.left = newNode
				break
			}
			node = node.left
		} else if value > node.data {
			// insert if right node is null
			if node.right == nil {
				node.right = newNode
				break
			}
			node = node.right
		} else {
			break
		}
	}
}

// Delete a value from binary search tree
func (st *SearchTree) Delete(value int) {
	// Find the node
	node := st.root
	var parent *Node
	for node != nil {
		parent = node
		if value < node.data {
			node = node.left
		} else if value > node.data {
			node = node.right
		} else {
			break
		}
	}
	if node == nil {
		return
	}

	// The node has two children
	if node.left != nil && node.right != nil {
		min := node.left
		minParent := node
		for min.left != nil {
			minParent, min = min, min.left
		}

		node.data = min.data          // Exchange min data to current node
		node, parent = min, minParent // min node to delete
	}

	// Delete root node
	if parent == nil {
		st.root = nil
	}

	// The node has one or none children
	var child *Node
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	// remove the node
	if parent.left == node {
		parent.left = child
	} else {
		parent.right = child
	}
}
