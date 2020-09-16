package stack

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

func GoWrap(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f()
	}()
}

func TestProducerAndConsumer(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	qu := NewRingQueue(3)
	producer := NewProducer(qu)
	consumer1 := NewConsumer("consumer 1", qu)
	consumer2 := NewConsumer("consumer 2", qu)
	ctx, cel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)
	GoWrap(func() {
		producer.Gen(ctx)
		wg.Done()
	})

	wg.Add(1)
	GoWrap(func() {
		consumer1.Exec(ctx)
		wg.Done()
	})

	wg.Add(1)
	GoWrap(func() {
		consumer2.Exec(ctx)
		wg.Done()
	})

	<-sigs
	cel()
	wg.Wait()
	return
}
