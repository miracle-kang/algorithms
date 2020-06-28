package tree_test

import (
	"testing"

	"../tree"
)

func TestBinarySearchTree(t *testing.T) {
	bsTree := tree.NewSearchTree()
	bsTree.Insert(33)
	bsTree.Insert(16)
	bsTree.Insert(13)
	bsTree.Insert(18)
	bsTree.Insert(15)
	bsTree.Insert(17)
	bsTree.Insert(25)
	bsTree.Insert(19)
	bsTree.Insert(27)
	bsTree.Insert(18)
	bsTree.Insert(50)
	bsTree.Insert(34)
	bsTree.Insert(58)
	bsTree.Insert(51)
	bsTree.Insert(66)
	bsTree.Insert(55)

}
