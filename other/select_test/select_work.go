package main

import (
	"fmt"
	"time"
)

var (
	stop chan int
)

func printHello() {
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				fmt.Printf("chan close")
				return
			}
		case <-time.After(time.Second * 3):
			fmt.Printf("run ...")
		}
	}
}

func main() {
	stop = make(chan int)
	go printHello()

	time.Sleep(time.Second * 9)
	close(stop)
	select {}
}
