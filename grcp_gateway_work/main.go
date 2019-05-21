package main

import (
	"fmt"
	"practiceDemo/grcp_gateway_work/sserver"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	addr := "127.0.0.1:32111"
	s := sserver.NewServerSimple(addr)
	s.Start()
}
