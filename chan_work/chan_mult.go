package chan_work

import "fmt"

type ChanWork struct {
	cache  chan interface{}
	stopCh chan int
}

func NewChanWork(num int64) *ChanWork {
	return &ChanWork{cache: make(chan interface{}, num), stopCh: make(chan int)}
}

func (c *ChanWork) Push(data interface{}) {
	c.cache <- data
}

func (c *ChanWork) Subscribe() <-chan interface{} {
	return c.cache
}

func (c *ChanWork) UseChan(flag string) {
	for {
		select {
		case dt, ok := <-c.Subscribe():
			if !ok {
				fmt.Println("UseChan error subscribe not ok")
				return
			}
			fmt.Println(flag, dt)
		case <-c.stopCh:
			return
		}
	}
}

func (c *ChanWork) Stop() {
	close(c.stopCh)
}
