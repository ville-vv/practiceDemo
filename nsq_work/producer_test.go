// @File     : producer_test
// @Author   : Ville
// @Time     : 19-9-16 下午5:57
// nsq_work
package nsq_work

import (
	"fmt"
	"testing"
	"time"
)

func TestProducer_PublishToJson(t *testing.T) {
	prd, err := NewProducer("127.0.0.1:4150")
	if err != nil {
		t.Error(err)
		return
	}
	stop := make(chan int)
	go func() {
		id := 0
		loop := true
		for loop {
			select {
			case <-stop:
				loop = false
			case <-time.After(time.Millisecond * 500):
				err = prd.PublishToJson("ville_nsq_test", map[string]interface{}{"ID": id, "name": "Name"})
				if err != nil {
					t.Error(err)
					return
				}
			}
			id++
		}
		fmt.Println("结束生产者")
	}()

	_, err = NewConsumer("ville_nsq_test", "ville_channel", "127.0.0.1:4150", func(msg []byte) error {
		fmt.Println("收到消息", string(msg))
		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 10)
	stop <- 0
	time.Sleep(time.Millisecond)
}
