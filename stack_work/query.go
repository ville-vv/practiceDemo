package stack

import "sync"

type Query struct {
	lock     sync.RWMutex
	head     *Node
	rear     *Node
	length   int32
	capacity int32
}

func NewQuery() *Query {
	front := &Node{
		value:    nil,
		previous: nil,
	}

	rear := &Node{
		value:    nil,
		previous: front,
	}

	front.Next = rear

	return &Query{
		head: front,
		rear: rear,
	}
}

// push 插入到新尾巴的前面一个节点
func (sel *Query) Push(val interface{}) {
	sel.lock.Lock()
	defer sel.lock.Unlock()

	node := newNode(val)
	// 头节点的下一个节点
	if sel.length == 0 {
		sel.head.Next = node
	}
	// 新节点前一个节点指向 尾节点的前一个节点
	node.previous = sel.rear.previous
	// 新节点指向尾节点
	node.Next = sel.rear

	// 尾节点的前一个节点的下一个节点指向  新的节点
	sel.rear.previous.Next = node
	// 尾节点之前一个节点指向 新的节点
	sel.rear.previous = node
	sel.length++
}

func (sel *Query) Shift() *Node {
	sel.lock.RLock()
	defer sel.lock.RUnlock()
	if sel.length == 0 {
		return nil
	}
	val := sel.head.Next
	sel.head.Next = val.Next
	val.Next = nil
	val.previous = nil
	sel.length--
	return val
}

func (sel *Query) Pop() *Node {
	sel.lock.RLock()
	defer sel.lock.RUnlock()
	val := sel.rear.previous
	sel.rear.previous = sel.rear.previous.previous
	sel.rear.previous.Next = sel.rear
	val.previous = nil
	val.Next = nil
	sel.length--
	return val
}

func (sel *Query) Length() int32 {
	return sel.length
}
