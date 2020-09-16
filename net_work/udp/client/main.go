package main

import (
	"fmt"
	"os"
	"os/signal"
	"practiceDemo/net_work/udp/svc"
	"syscall"
	"time"
)

func main() {
	//
	cli := svc.NewUdpClient("192.168.2.13", "19999")
	_ = cli.SetLocalAddr("0.0.0.0:20000")
	go func() {
		for {
			cli.Recv()
		}
	}()

	err := cli.Send()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	time.Sleep(time.Second * 1)
	//svc.RecBoardCast()
}
