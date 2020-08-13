package main

import (
	"fmt"
	"practiceDemo/nsq_work/svce"
	"time"
)

func main() {
	prd, err := svce.NewProducer("127.0.0.1:4150")
	if err != nil {
		fmt.Println("生产者链接错误", err)
		return
	}

	for {
		err = prd.PublishString("topic01", fmt.Sprintf("发送一个消息%s", time.Now().String()))
		if err != nil {
			fmt.Println("生产者发送想消息错误", err)
			return
		}
		time.Sleep(time.Second)
	}
}
