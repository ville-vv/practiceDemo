package stack

import (
	"fmt"
	"testing"
	"time"
)

func TestStackSlice_Push(t *testing.T) {
	s := NewStackSlice(10100)
	s.Push(30)
	s.Push(20)
	s.Push(40)
	s.Push(50)
	s.Push(10)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println("---------------------------")
	s.Push(30)
	s.Push(20)
	s.Push(40)
	s.Push(50)
	s.Push(10)

	fmt.Println(s.Shift())
	fmt.Println(s.Shift())
	fmt.Println(s.Shift())
	fmt.Println(s.Shift())
	fmt.Println(s.Shift())
	fmt.Println(s.Shift())
}

func TestStackSlice_Push2(t *testing.T) {
	num := 100000000
	po := NewStackSlice(0)
	start := time.Now().UnixNano()
	for i := 0; i < num; i++ {
		po.Push(i)
	}
	end := time.Now().UnixNano()

	fmt.Println("入栈时间：", (end-start)/1e6)
	start = time.Now().UnixNano()
	for i := 0; i < num; i++ {
		//fmt.Println(po.Shift())
		po.Pop()
	}
	end = time.Now().UnixNano()
	fmt.Println("出栈时间：", (end-start)/1e6)
}

func TestStackSlice_Push3(t *testing.T) {
	po := NewStackSlice(100)
	po.Push(Do(sub))
	po.Push(Do(sum))
	po.Push(Do(sub))
	fmt.Println((po.Pop()).(Do)(30, 40))
	fmt.Println((po.Pop()).(Do)(30, 40))
	fmt.Println((po.Pop()).(Do)(50, 40))
}
