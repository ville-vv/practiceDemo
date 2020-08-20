package main

import (
	"os"
	"os/signal"
	"practiceDemo/net_work/udp/svc"
	"syscall"
	"time"
)

func main() {
	rd, _ := svc.NewUdpServer("0.0.0.0", "19999")
	//rd.AddRemoteAddr("172.16.5.48:20000")
	go func() {
		for {
			rd.ReadFromUdp()
			_ = rd.Send()
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	time.Sleep(time.Second * 1)
}
