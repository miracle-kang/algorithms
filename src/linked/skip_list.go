package linked

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	skipListP = 0.5
	maxLevel  = 16
)

// randLevel Generate a rand level in O(1)
func randLevel() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	level := 1
	for level < maxLevel && r.Float64() > skipListP {
		level++
	}
	return level
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// SkipNode is a skip list node
type SkipNode struct {
	// extends linked list node
	next *SkipNode
	// down point to
	down *SkipNode

	data int
}

// Value of the skip node
func (sn *SkipNode) Value() int {
	return sn.data
}

// SkipList is a list implement by Linked list
type SkipList struct {
	// extends from linked list
	heads []*SkipNode

	size       int
	levelCount int
}

// NewSkipList new and init a skip list
func NewSkipList() *SkipList {
	heads := make([]*SkipNode, maxLevel)

	heads[0] = &SkipNode{}
	for i := 1; i < maxLevel; i++ {
		heads[i] = &SkipNode{down: heads[i-1]}
	}
	return &SkipList{heads: heads, levelCount: 1}
}

// Size of skip list
func (sl *SkipList) Size() int {
	return sl.size
}

// Find a value from skip list
func (sl *SkipList) Find(value int) *SkipNode {
	node := sl.heads[sl.levelCount-1]
	for node.down != nil || (node.next != nil && node.next.data < value) {
		if node.next != nil && node.next.data < value {
			node = node.next
		} else {
			node = node.down
		}
	}

	if node.next != nil && node.next.data == value {
		return node.next
	}
	return nil
}

// Insert a value to skip list
func (sl *SkipList) Insert(value int) {
	level := randLevel()

	// new nodes
	newNodes := make([]*SkipNode, level)
	newNodes[0] = &SkipNode{data: value}
	for i := 1; i < level; i++ {
		newNodes[i] = &SkipNode{data: value, down: newNodes[i-1]}
	}

	i := max(sl.levelCount-1, level-1)
	node := sl.heads[i]
	for i >= 0 {
		if node.next != nil && node.next.data < value {
			node = node.next
			continue
		} else {
			if i < level {
				newNodes[i].next = node.next
				node.next = newNodes[i]
			}
			if node.down != nil {
				node = node.down
			}
			i--
		}
	}

	sl.size++
	if sl.levelCount < level {
		sl.levelCount = level
	}
}

// Delete a value from skip list
func (sl *SkipList) Delete(value int) bool {
	backs := make([]*SkipNode, sl.levelCount)

	level := sl.levelCount - 1
	node := sl.heads[sl.levelCount-1]

	for level >= 0 {
		if node.next != nil && node.next.data < value {
			node = node.next
		} else {
			backs[level] = node

			if node.down != nil {
				node = node.down
			}
			level--
		}
	}

	if node.next == nil || node.next.data != value {
		return false
	}

	// delete cache
	for i := sl.levelCount - 1; i >= 0; i-- {
		if backs[i].next != nil && backs[i].next.data == value {
			backs[i].next = backs[i].next.next
		}
	}

	sl.resizeLevelCount()
	sl.size--
	return true
}

func (sl *SkipList) resizeLevelCount() {
	level := sl.levelCount - 1
	for ; level >= 0 && sl.heads[level].next == nil; level-- {
		sl.levelCount--
	}
}

// Print all node
func (sl *SkipList) Print() {
	for i := sl.levelCount - 1; i >= 0; i-- {
		node := sl.heads[i]

		fmt.Print("head->")
		for node.next != nil {
			node = node.next
			fmt.Print(node.data, "->")
		}
		fmt.Println("tail")
	}
}
