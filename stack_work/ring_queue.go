// 实现一个环形队列，使用一个生产者两个消费者消费，消费者速度消费要慢于生产者
package stack

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type RingQueue struct {
	list   []interface{}
	llen   int64
	ccap   int64
	pshIdx int64
	shtIdx int64
}

func NewRingQueue(ccap int64) *RingQueue {
	return &RingQueue{ccap: ccap, list: make([]interface{}, ccap)}
}

func (r *RingQueue) Push(val interface{}) error {
	if atomic.LoadInt64(&r.llen) >= r.ccap {
		return errors.New("ring queue have full")
	}

	r.list[r.pshIdx] = val

	atomic.AddInt64(&r.pshIdx, 1)
	atomic.AddInt64(&r.llen, 1)
	if atomic.LoadInt64(&r.pshIdx) >= r.ccap {
		atomic.StoreInt64(&r.pshIdx, 0)
	}
	return nil
}

func (r *RingQueue) Shift() interface{} {
	if atomic.LoadInt64(&r.llen) == 0 {
		return nil
	}
	temp := r.list[r.shtIdx]
	atomic.AddInt64(&r.llen, -1)
	atomic.AddInt64(&r.shtIdx, 1)
	if atomic.LoadInt64(&r.shtIdx) >= r.ccap {
		atomic.StoreInt64(&r.shtIdx, 0)
	}
	return temp
}

type IQueue interface {
	Push(val interface{}) error
	Shift() interface{}
}

type Producer struct {
	qu IQueue
}

func (sel *Producer) Gen(ctx context.Context) {
	tm := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("安全退出")
			return
		case <-tm.C:
			dt := rand.Int63n(100)
			if err := sel.qu.Push(dt); err != nil {
				fmt.Println("生产者错误")
				return
			}
			fmt.Printf("生产者生产:%d\n", dt)
		}
	}
}

func NewProducer(qu IQueue) *Producer {
	return &Producer{qu: qu}
}

type Consumer struct {
	name string
	qu   IQueue
}

func NewConsumer(name string, qu IQueue) *Consumer {
	return &Consumer{
		name: name,
		qu:   qu,
	}
}

func (sel *Consumer) Exec(ctx context.Context) {
	for {
		tm := time.NewTicker(time.Second * 3)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("安全退出")
				return
			case <-tm.C:
				fmt.Printf("消费者[%s]消费内容：%v \n", sel.name, sel.qu.Shift())
			}
		}
	}
}
