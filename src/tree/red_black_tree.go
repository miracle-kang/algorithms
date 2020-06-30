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

// Size of the red-black tree
func (rbt *RBTree) Size() int {
	return rbt.size
}

// Sorted the binary search tree to an array
func (rbt *RBTree) Sorted() []int {
	arr := make([]int, rbt.size)
	index := 0
	innerTraversalRBTree(rbt.root, arr, &index)

	return arr
}

func innerTraversalRBTree(root *RBNode, arr []int, index *int) {
	if root == nil {
		return
	}

	innerTraversalRBTree(root.left, arr, index)
	arr[*index] = root.data
	*index++
	innerTraversalRBTree(root.right, arr, index)
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
	insertRepairTree(newNode)

	// Find the new root
	root := newNode
	for root.parent != nil {
		root = root.parent
	}
	rbt.root = root
	rbt.size++
}

// Repair the tree, return the new root
// 如果插入节点的父节点是黑色的，那我们什么都不用做，它仍然满足红黑树的定义。
// 如果插入的节点是根节点，那我们直接改变它的颜色，把它变成黑色就可以了。
// 其它情况则需要做下面的调整
func insertRepairTree(node *RBNode) {
	if parent(node) == nil {
		node.color = BLACK
	} else if parent(node).color == BLACK {
		// Nothing to do
	} else if uncle(node) != nil {
		if uncle(node).color == RED {
			insertCase1(node)
		} else if parent(node).right == node {
			insertCase2(node)
		} else if parent(node).left == node {
			insertCase3(node)
		}
	}
}

// CASE 1：如果关注节点是 a，它的叔叔节点 d 是红色，我们就依次执行下面的操作：
// 将关注节点 a 的父节点 b、叔叔节点 d 的颜色都设置成黑色；
// 将关注节点 a 的祖父节点 c 的颜色设置成红色；
// 关注节点变成 a 的祖父节点 c；
// 跳到 CASE 2 或者 CASE 3。
func insertCase1(node *RBNode) {
	parent(node).color = BLACK
	uncle(node).color = BLACK
	grandParent(node).color = RED

	insertRepairTree(grandParent(node))
}

// CASE 2：如果关注节点是 a，它的叔叔节点 d 是黑色，关注节点 a 是其父节点 b 的右子节点，我们就依次执行下面的操作：
// 关注节点变成节点 a 的父节点 b；
// 围绕新的关注节点b 左旋；
// 跳到 CASE 3。
func insertCase2(node *RBNode) {
	node = parent(node)
	rotateLeft(node)
	insertCase3(node)
}

// CASE 3：如果关注节点是 a，它的叔叔节点 d 是黑色，关注节点 a 是其父节点 b 的左子节点，我们就依次执行下面的操作：
// 围绕关注节点 a 的祖父节点 c 右旋；
// 将关注节点 a 的父节点 b、兄弟节点 c 的颜色互换。
// 调整结束。
func insertCase3(node *RBNode) {
	rotateRight(grandParent(node))
	parent(node).color = BLACK
	sibling(node).color = RED
}

// Delete a value from red-black tree
func (rbt *RBTree) Delete(value int) {

}
