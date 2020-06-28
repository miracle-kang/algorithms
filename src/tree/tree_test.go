package tree_test

import (
	"fmt"
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
	bsTree.Insert(50)
	bsTree.Insert(34)
	bsTree.Insert(58)
	bsTree.Insert(51)
	bsTree.Insert(66)
	bsTree.Insert(55)
	// insert case
	tree.LayerPrint(bsTree.Root())
	if bsTree.Size() != 15 {
		t.Error("Unexpected tree size ", bsTree.Size(), ", expect 15")
	}

	// find case
	node := bsTree.Find(17)
	if node == nil || node.Value() != 17 {
		t.Error("Unexpected value ", node.Value(), ", expect 17")
	}
	node = bsTree.Find(14)
	if node != nil {
		t.Error("Unexpected result ", node.Value(), ", expect 14 is not exists")
	}

	// delete case
	if !bsTree.Delete(13) {
		t.Error("Unexpected delete 13 result, expect true")
	}
	if bsTree.Delete(14) {
		t.Error("Unexpected delete 14 result, expect false")
	}
	bsTree.Delete(18)
	bsTree.Delete(55)

	if bsTree.Size() != 12 {
		t.Error("Unexpected tree size ", bsTree.Size(), ", expect 12")
	}
	node = bsTree.Find(18)
	if node != nil {
		t.Error("Unexpected result ", node.Value(), ", expect 18 is deleted")
	}
	tree.LayerPrint(bsTree.Root())

	sortedArr := bsTree.Sorted()
	fmt.Println(sortedArr)
}
