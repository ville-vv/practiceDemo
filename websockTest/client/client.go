package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
)

var origin = "http://127.0.0.1:9090/"
var url = "ws://127.0.0.1:9090/echo"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close()//关闭连接
}
