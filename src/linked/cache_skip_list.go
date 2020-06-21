package linked

import "fmt"

// Already define in skip_list.go
// const (
// 	skipListP = 0.5
// 	maxLevel  = 16 // Keep this in 2^n, help to generate a rand level
// )

// CacheNode is a cache skip list node
type CacheNode struct {
	// data
	data int

	// Cache nodes
	cacheNodes []*CacheNode
}

// Value of the node
func (cn *CacheNode) Value() int {
	return cn.data
}

// CacheSkipList is a list implement by skip cache
type CacheSkipList struct {
	head       *CacheNode
	size       int
	levelCount int
}

// NewCacheSkipList : new and init a CacheSkipList
func NewCacheSkipList() *CacheSkipList {
	head := &CacheNode{cacheNodes: make([]*CacheNode, maxLevel)}
	return &CacheSkipList{head: head, size: 0, levelCount: 1}
}

// LevelCount for the cache skip list
func (csl CacheSkipList) LevelCount() int {
	return csl.levelCount
}

// Find a value from skip list
func (csl *CacheSkipList) Find(value int) *CacheNode {
	node := csl.head
	for i := csl.levelCount - 1; i >= 0; i-- {
		// Get node by index level
		for node.cacheNodes[i] != nil && node.cacheNodes[i].data < value {
			// forward at index level
			node = node.cacheNodes[i]
		}
	}
	// continue forward a index at level 0
	node = node.cacheNodes[0]

	// current at level 0, and the node value is >= by value
	if node.data == value {
		return node
	}
	return nil
}

// Insert a value to skip list
func (csl *CacheSkipList) Insert(value int) {
	// Already define in skip_list.go
	level := randLevel()
	newNode := &CacheNode{data: value, cacheNodes: make([]*CacheNode, level)}

	node := csl.head
	for i := max(csl.levelCount-1, level); i >= 0; i-- {
		// Find the index that insert node to
		for node.cacheNodes[i] != nil && node.cacheNodes[i].data < value {
			node = node.cacheNodes[i]
		}

		if i < level {
			// Insert the new node at level i
			newNode.cacheNodes[i] = node.cacheNodes[i]
			node.cacheNodes[i] = newNode
		}
	}

	// update levelCount
	if csl.levelCount < level {
		csl.levelCount = level
	}
	csl.size++
}

// Delete a value from skip list
func (csl *CacheSkipList) Delete(value int) bool {
	backs := make([]*CacheNode, csl.levelCount)
	node := csl.head
	for i := csl.levelCount - 1; i >= 0; i-- {
		for node.cacheNodes[i] != nil && node.cacheNodes[i].data < value {
			node = node.cacheNodes[i]
		}

		// record back nodes
		backs[i] = node
	}

	// if next node value not equals than value
	if node.cacheNodes[0] == nil || node.cacheNodes[0].data != value {
		return false
	}

	// remove the node and it's index
	for i := csl.levelCount - 1; i >= 0; i-- {
		if backs[i].cacheNodes[i] != nil && backs[i].cacheNodes[i].data == value {
			backs[i].cacheNodes[i] = backs[i].cacheNodes[i].cacheNodes[i]
		}
	}
	csl.resizeLevelCount()
	csl.size--
	return true
}

// Size of cache skip list
func (csl *CacheSkipList) Size() int {
	return csl.size
}

// Resize level count
func (csl *CacheSkipList) resizeLevelCount() {
	head := csl.head
	for csl.levelCount > 1 && head.cacheNodes[csl.levelCount-1] == nil {
		csl.levelCount--
	}
}

// Print the skip list
func (csl *CacheSkipList) Print() {
	for i := csl.levelCount - 1; i >= 0; i-- {
		node := csl.head
		fmt.Print("Level ", i, ": head->")
		for node.cacheNodes[i] != nil {
			node = node.cacheNodes[i]
			print(node.data, "->")
		}
		fmt.Println("tail")
	}
}
