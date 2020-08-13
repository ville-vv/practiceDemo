package main

import (
	"fmt"
	"practiceDemo/nsq_work/svce"
)

func main() {
	go func() {
		svce.NewConsumer("topic01", "channel_01", "127.0.0.1:4150", func(msg []byte) error {
			fmt.Println("第一个消费者：", string(msg))
			return nil
		})
	}()

	//go func() {
	//	svce.NewConsumer("topic01", "channel_01", "127.0.0.1:4150", func(msg []byte) error {
	//		fmt.Println("第二个消费者：", string(msg))
	//		return nil
	//	})
	//}()

	select {}
}
