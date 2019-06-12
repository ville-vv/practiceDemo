package stack

import "sync"

type ItemNode []interface{}

type StackSlice struct {
	items    ItemNode
	length   int
	capacity int
	sync.RWMutex
}

func NewStackSlice(cp int) *StackSlice {
	return &StackSlice{
		items:    make(ItemNode, 0, cp),
		capacity: cp,
	}
}

func (sel *StackSlice) Pop() interface{} {
	sel.Lock()
	defer sel.Unlock()
	length := len(sel.items)
	if length == 0 {
		return nil
	}
	item := sel.items[length-1]
	sel.items = sel.items[:length-1]
	return item
}

func (sel *StackSlice) Push(val interface{}) {
	sel.Lock()
	defer sel.Unlock()
	sel.items = append(sel.items, val)
}

func (sel *StackSlice) Shift() interface{} {
	sel.Lock()
	defer sel.Unlock()
	length := len(sel.items)
	if length == 0 {
		return nil
	}
	item := sel.items[0]
	if length > 1 {
		sel.items = sel.items[1:length]
	} else {
		sel.items = make([]interface{}, 0, sel.capacity)
	}
	return item
}
