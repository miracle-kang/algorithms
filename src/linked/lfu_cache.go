package linked

type LinkNode struct {
	key, val, cnt int
	prev, next    *LinkNode
}

func NewLinkNode(key, val int) *LinkNode {
	return &LinkNode{
		key: key,
		val: val,
		cnt: 1,
	}
}

type CountLinked struct {
	size       int
	head, tail *LinkNode
}

func NewCountLinked(node *LinkNode) *CountLinked {
	return &CountLinked{
		size: 1,
		head: node,
		tail: node,
	}
}

func (this *CountLinked) addFirst(node *LinkNode) {
	node.next = this.head
	this.head.prev = node
	this.head = node
	this.size++
}

func (this *CountLinked) removeLast() *LinkNode {
	node := this.tail
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
	}
	this.size--
	return node
}

func (this *CountLinked) remove(node *LinkNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}
	node.prev = nil
	node.next = nil
	this.size--
}

type LFUCache struct {
	cache          map[int]*LinkNode    // key -> *LinkNode
	cntCache       map[int]*CountLinked // count -> *CountLinked
	size, capacity int
	minCnt         int
}

func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*LinkNode, capacity+1),
		cntCache: make(map[int]*CountLinked, capacity+1),
		size:     0,
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}

	cnt := node.cnt
	if this.cntCache[cnt].size == 1 {
		delete(this.cntCache, cnt)
		if cnt == this.minCnt {
			this.minCnt++
		}
	} else {
		this.cntCache[cnt].remove(node)
	}

	node.cnt++
	if cntLink, ok := this.cntCache[node.cnt]; ok {
		cntLink.addFirst(node)
	} else {
		this.cntCache[node.cnt] = NewCountLinked(node)
	}
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.cache[key].val = value
		return
	}

	// Remove min cnt node
	if this.size >= this.capacity {
		minNode := this.cntCache[this.minCnt]
		node := minNode.removeLast()
		if minNode.size == 0 {
			delete(this.cntCache, this.minCnt)
		}
		delete(this.cache, node.key)
		this.size--
	}

	// Add new node
	node := NewLinkNode(key, value)
	this.cache[key] = node
	if cntNode, ok := this.cntCache[node.cnt]; ok {
		cntNode.addFirst(node)
	} else {
		this.cntCache[node.cnt] = NewCountLinked(node)
	}
	this.size++
	this.minCnt = node.cnt
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
