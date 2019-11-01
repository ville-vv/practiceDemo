package stack

import (
	"errors"
	"sync"
)

func newNode(val interface{}) *node {
	return &node{value: val}
}

func New() Queue {
	return newQueue()
}

type Queue interface {
	Pop() *node
	Push(*node) error
	Shift() *node
	Length() int64
}

type queue struct {
	lock     sync.Mutex
	head     *node
	rear     *node
	length   int64
	capacity int64
}

func newQueue(args ...interface{}) *queue {
	max := int64(100000000)
	if len(args) > 0 {
		switch args[0].(type) {
		case int64:
			max = args[0].(int64)
		case int:
			max = int64(args[0].(int))
		}
	}
	return &queue{
		capacity: max,
	}
}
func (sel *queue) Pop() *node {
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
func (sel *queue) Shift() *node {
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
func (sel *queue) Push(n *node) error {
	if sel.length >= sel.capacity {
		return errors.New("over max num for stack")
	}
	sel.push(n)
	return nil
}
func (sel *queue) push(top *node) {
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
func (sel *queue) Length() int64 {
	return sel.length
}

var (
	defaultQueue = newQueue()
)

func Pop() *node {
	return defaultQueue.Pop()
}
func Push(v interface{}) error {
	return defaultQueue.Push(&node{value: v})
}
func Shift() *node {
	return defaultQueue.Shift()
}
func Length() int64 {
	return defaultQueue.Length()
}
