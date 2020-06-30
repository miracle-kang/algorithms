package tree

// BSNode is a binary search tree node
type BSNode struct {
	left, right *BSNode
	data        int
}

// Value of the node
func (bsn *BSNode) Value() int {
	return bsn.data
}

// Left of the node child
func (bsn *BSNode) Left() INode {
	return bsn.left
}

// Right of the node child
func (bsn *BSNode) Right() INode {
	return bsn.right
}

// BSTree a binary search tree implemention
type BSTree struct {
	root *BSNode
	size int
}

// Size of binary search tree
func (bst *BSTree) Size() int {
	return bst.size
}

// Root of binary search tree
func (bst *BSTree) Root() *BSNode {
	return bst.root
}

// NewBSTree new and initialize a binary search tree
func NewBSTree() *BSTree {
	return &BSTree{}
}

// Find a value from binary search tree
func (bst *BSTree) Find(value int) *BSNode {

	node := bst.root
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
func (bst *BSTree) Insert(value int) {
	newNode := &BSNode{data: value}
	if bst.root == nil {
		bst.root = newNode
		bst.size = 1
		return
	}

	node := bst.root
	for node != nil {
		if value < node.data {
			// insert if left node is null
			if node.left == nil {
				node.left = newNode
				bst.size++
				break
			}
			node = node.left
		} else if value > node.data {
			// insert if right node is null
			if node.right == nil {
				node.right = newNode
				bst.size++
				break
			}
			node = node.right
		} else {
			break
		}
	}
}

// Delete a value from binary search tree
func (bst *BSTree) Delete(value int) bool {
	// Find the node
	node := bst.root
	var parent *BSNode
	for node != nil && node.data != value {
		parent = node
		if value < node.data {
			node = node.left
		} else if value > node.data {
			node = node.right
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
	var child *BSNode
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	// remove the node
	if parent == nil {
		bst.root = nil
	} else if parent.left == node {
		parent.left = child
	} else {
		parent.right = child
	}
	bst.size--
	return true
}

// Sorted the binary search tree to an array
func (bst *BSTree) Sorted() []int {
	arr := make([]int, bst.size)
	index := 0
	innerTraversalBSTree(bst.root, arr, &index)

	return arr
}

func innerTraversalBSTree(root *BSNode, arr []int, index *int) {
	if root == nil {
		return
	}

	innerTraversalBSTree(root.left, arr, index)
	arr[*index] = root.data
	*index++
	innerTraversalBSTree(root.right, arr, index)
}
