package stack

import (
	"fmt"
	"testing"
	"time"
)

func TestQuery_Push(t *testing.T) {
	q := New()
	q.Push(&node{value: 30})
	q.Push(&node{value: 50})
	q.Push(&node{value: 90})
	q.Push(&node{value: 100})
	q.Push(&node{value: 5})

	fmt.Println(q.Shift().Value())
	fmt.Println(q.Shift().Value())
	fmt.Println(q.Shift().Value())
	fmt.Println(q.Shift().Value())
	fmt.Println(q.Shift().Value())

	q.Push(&node{value: 30})
	q.Push(&node{value: 50})
	q.Push(&node{value: 90})
	q.Push(&node{value: 100})
	q.Push(&node{value: 5})

	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
}

func TestQuery_Push2(t *testing.T) {
	po := New()
	num := 10000000
	start := time.Now().UnixNano()
	for i := 0; i < num; i++ {
		po.Push(&node{value: i})
	}
	end := time.Now().UnixNano()
	fmt.Println("Push的时间：", (end-start)/1e6)

	fmt.Println("长度：", po.Length())

	start = time.Now().UnixNano()
	for i := 0; i < num; i++ {
		po.Pop()
	}
	end = time.Now().UnixNano()
	fmt.Println("Pop的时间：", (end-start)/1e6)
}

type Do func(a, b int) int

func sub(a, b int) int {
	return a - b
}

func sum(a, b int) int {
	return a + b
}

var po = New()
var poSlice = NewStackSlice(100000000)

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		po.Push(&node{value: i})
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		po.Push(&node{value: i})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		po.Pop()
	}
}

func BenchmarkStackSlice_Pop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		poSlice.Push(i)
	}
}

func BenchmarkStackSlice_Push(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		poSlice.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		poSlice.Pop()
	}
}
