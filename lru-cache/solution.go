package lru_cache

type LRUCache struct {
	cap  int
	m    map[int]*node
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 || capacity > 3000 {
		panic("cap should be greater than 0")
	}

	return LRUCache{
		cap: capacity,
		m:   make(map[int]*node, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	validateKey(key)
	if len(c.m) == 0 {
		return -1
	}

	e, ok := c.m[key]
	if !ok {
		return -1
	}

	c.makeHead(e)
	return e.value
}

func (c *LRUCache) Put(key int, value int) {
	validateKey(key)
	validateValue(value)

	// Special, short circuit case of cap 1. We dont use linked list in this case
	if c.cap == 1 {
		// Initial case
		if len(c.m) == 0 {
			e := &node{
				key:   key,
				value: value,
			}
			c.head = e
			c.m[key] = e
			return
		}

		// Just updating node value in case of requested key is already in cache
		e, ok := c.m[key]
		if ok {
			e.value = value
			return
		}

		delete(c.m, c.head.key)
		newE := &node{
			key:   key,
			value: value,
		}
		c.head = newE
		c.m[key] = newE
		return
	}

	// General case.
	// ====================
	// Initial case
	if len(c.m) == 0 {
		e := &node{
			key:   key,
			value: value,
		}
		c.head = e
		c.tail = e
		c.m[key] = e
		return
	}

	// After initial case it is guaranteed to have head and tail nodes set
	//
	// In case requested key is already in cache we updating node value and
	// promoting the node to head of the linked list
	if e, ok := c.m[key]; ok {
		e.value = value
		c.makeHead(e)
		return
	}

	// If cache is full we need to remove the least recently used node
	if c.cap == len(c.m) {
		delete(c.m, c.tail.key)
		c.tail = c.tail.prev
	}

	e := &node{
		key:   key,
		value: value,
		next:  c.head,
	}
	c.head.prev = e
	c.head = e
	c.m[key] = e
}

func (c *LRUCache) makeHead(e *node) {
	if c.head == e {
		return
	}

	if c.tail == e {
		c.tail = e.prev
	}

	h := c.head
	e1 := e.prev
	e2 := e
	e3 := e.next

	if e1 != nil {
		e1.next = e3
	}

	if e3 != nil {
		e3.prev = e1
	}

	h.prev = e2
	e2.next = h
	e2.prev = nil
	c.head = e2
}

func validateKey(key int) {
	if key < 0 || key > 10000 {
		panic("key should be in range [1, 10000]")
	}
}

func validateValue(value int) {
	if value < 0 || value > 100000 {
		panic("value should be in range [0, 10000]")
	}
}

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}
