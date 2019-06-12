package stack

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type Node struct {
	value    interface{}
	previous *Node
	Next     *Node
}

func (sel *Node) ToString() string {
	return ""
}
func (sel *Node) ToInt() int {
	d := 0
	switch sel.value.(type) {
	case string:
		d, _ := strconv.Atoi(sel.value.(string))
		return d
	case int64:
		d = int(sel.value.(int64))
	case int32:
		d = int(sel.value.(int32))
	case int:
		d = sel.value.(int)
	}
	return d
}
func (sel *Node) ToInt64() int64 {
	d := int64(0)
	switch sel.value.(type) {
	case string:
		d, _ := strconv.ParseInt(sel.value.(string), 10, 64)
		return d
	case int64:
		d = sel.value.(int64)
	case int32:
		d = int64(sel.value.(int32))
	case int:
		d = int64(sel.value.(int))
	}
	return d
}
func (sel *Node) ToInt32() int {
	return 0
}
func (sel *Node) Value() interface{} {
	return sel.value
}

func newNode(val interface{}) *Node {
	return &Node{value: val}
}

type Stack struct {
	sync.RWMutex
	head   *Node
	rear   *Node
	length int32
	cache  chan interface{}
}

func NewPool() *Stack {
	return &Stack{}
}
func NewPoolChan() *Stack {
	s := &Stack{
		cache: make(chan interface{}, 1),
	}
	//go s.readChan()
	return s
}

func (sel *Stack) Pop() *Node {
	sel.RLock()
	defer sel.RUnlock()
	val := &Node{
		value: sel.head.value,
	}
	sel.head = sel.head.Next
	atomic.AddInt32(&sel.length, -1)
	return val
}

func (sel *Stack) Push(v interface{}) {
	sel.Lock()
	defer sel.Unlock()
	sel.push(v)
}
func (sel *Stack) push(v interface{}) {
	top := newNode(v)
	if sel.length == 0 {
		sel.head = top
		sel.rear = sel.head
	}

	//head := sel.head
	top.Next = sel.head
	sel.head.previous = top
	sel.head = top
	atomic.AddInt32(&sel.length, 1)
}

func (sel *Stack) Shift() *Node {
	val := sel.rear
	sel.rear = sel.rear.previous
	sel.rear.Next = nil
	val.Next = nil
	val.previous = nil
	return val
}

func (sel *Stack) Length() int32 {
	return sel.length
}

func (sel *Stack) readChan() {
	for v := range sel.cache {
		//fmt.Println(v)
		sel.push(v)
	}
}

func (sel *Stack) PushChan(val interface{}) {
	sel.cache <- val
	sel.push(<-sel.cache)
}
