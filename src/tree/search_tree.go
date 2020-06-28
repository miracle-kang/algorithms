package tree

// SearchTree a binary search tree implemention
type SearchTree struct {
	root *Node
	size int
}

// Size of binary search tree
func (st *SearchTree) Size() int {
	return st.size
}

// Root of binary search tree
func (st *SearchTree) Root() *Node {
	return st.root
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
		st.size = 1
		return
	}

	node := st.root
	for node != nil {
		if value < node.data {
			// insert if left node is null
			if node.left == nil {
				node.left = newNode
				st.size++
				break
			}
			node = node.left
		} else if value > node.data {
			// insert if right node is null
			if node.right == nil {
				node.right = newNode
				st.size++
				break
			}
			node = node.right
		} else {
			break
		}
	}
}

// Delete a value from binary search tree
func (st *SearchTree) Delete(value int) bool {
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
		return false
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

	// The node has one or none children
	var child *Node
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	// remove the node
	if parent == nil {
		st.root = nil
	} else if parent.left == node {
		parent.left = child
	} else {
		parent.right = child
	}
	st.size--
	return true
}
