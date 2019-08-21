package stack

import (
	"errors"
	"sync"
)

func New() Queue {
	return NewNodeQueue()
}

type Queue interface {
	Pop() *Node
	Push(interface{}) error
	Shift() *Node
	Length() int64
}

type NodeQueue struct {
	lock     sync.Mutex
	head     *Node
	rear     *Node
	length   int64
	capacity int64
}

func NewNodeQueue(args ...interface{}) *NodeQueue {
	max := int64(100000000)
	if len(args) > 0 {
		switch args[0].(type) {
		case int64:
			max = args[0].(int64)
		case int:
			max = int64(args[0].(int))
		}
	}
	return &NodeQueue{
		capacity: max,
	}
}

func (sel *NodeQueue) Pop() *Node {
	sel.lock.Lock()
	defer sel.lock.Unlock()
	if sel.length <= 0 {
		sel.length = 0
		return nil
	}
	val := sel.head
	sel.head = sel.head.Next
	sel.length--
	val.previous = nil
	val.Next = nil
	return val
}
func (sel *NodeQueue) Shift() *Node {
	sel.lock.Lock()
	defer sel.lock.Unlock()

	if sel.length <= 0 {
		sel.length = 0
		return nil
	}

	val := sel.rear
	if sel.rear.previous == nil {
		sel.rear = sel.head
	} else {
		sel.rear = sel.rear.previous
		sel.rear.Next = nil
	}
	val.previous = nil
	val.Next = nil
	sel.length--
	return val
}
func (sel *NodeQueue) Push(v interface{}) error {
	if sel.length >= sel.capacity {
		return errors.New("over max num for stack")
	}
	sel.push(&Node{value: v})
	return nil
}
func (sel *NodeQueue) push(top *Node) {
	sel.lock.Lock()
	defer sel.lock.Unlock()
	if 0 == sel.length {
		sel.head = top
		sel.rear = sel.head
	}
	top.Next = sel.head
	sel.head.previous = top
	sel.head = top
	sel.length++
	return
}
func (sel *NodeQueue) Length() int64 {
	return sel.length
}
func (sel *NodeQueue) ResetPush(nd *Node) error {
	if sel.length >= sel.capacity {
		return errors.New("over max num for stack")
	}
	sel.push(nd)
	return nil
}
