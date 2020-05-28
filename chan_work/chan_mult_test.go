package chan_work

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 当有一个chan被多次订阅，chan只选择一个使用
func TestChanWork_UseChan(t *testing.T) {
	cw := NewChanWork(1000)
	go cw.UseChan("第一个：")
	go cw.UseChan("第二个：")
	go cw.UseChan("第三个：")
	go cw.UseChan("第四个：")
	go cw.UseChan("第五个：")
	go cw.UseChan("第六个：")
	cw.Push("weogijoisuio")
	cw.Push(map[string]interface{}{
		"object": "",
		"struct": 90,
	})
	cw.Push(93487985437)
	time.Sleep(time.Second)
}

func TestChanWork_Subscribe(t *testing.T) {
	cw := NewChanWork(10)
	ch := make(chan int, 1)
	cw.Push("Chan Message after subscribe")
	ch <- 0
	select {
	case <-cw.Subscribe():
		fmt.Println("chan work")
	case <-ch:
		fmt.Println("int chan")
	}
}

// 这个测试会出现错误
func TestChanWork_Subscribe1(t *testing.T) {
	cw := NewChanWork(10)
	// 没有 传值的 chan 会报错
	select {
	case <-cw.Subscribe():
		fmt.Println("chan work")
	}
}

func TestChanWork_OSName(t *testing.T) {
	fmt.Println(os.Hostname())
}
