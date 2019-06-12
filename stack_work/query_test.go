package stack

import (
	"fmt"
	"testing"
	"time"
)

func TestQuery_Push(t *testing.T) {
	q := NewQuery()
	q.Push(30)
	q.Push(50)
	q.Push(90)
	q.Push(3)
	q.Push(5)

	fmt.Println(q.Shift())
	fmt.Println(q.Shift())
	fmt.Println(q.Shift())
	fmt.Println(q.Shift())
	fmt.Println(q.Shift())

	q.Push(30)
	q.Push(50)
	q.Push(90)
	q.Push(3)
	q.Push(5)

	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
	fmt.Println(q.Pop().ToInt())
}

func TestQuery_Push2(t *testing.T) {
	po := NewQuery()
	num := 10000000
	start := time.Now().UnixNano()
	for i := 0; i < num; i++ {
		po.Push(i)
	}
	end := time.Now().UnixNano()
	fmt.Println("时间1：", (end-start)/1e6)

	fmt.Println("长度：", po.Length())

	start = time.Now().UnixNano()
	for i := 0; i < num; i++ {
		po.Pop()
	}
	end = time.Now().UnixNano()
	fmt.Println("时间2：", (end-start)/1e6)
}
