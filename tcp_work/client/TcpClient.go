package main

import (
	"time"
	"fmt"
	"io"
	"net"
)

type ServerConf struct{
	Host string
}

type TcpClient struct{
	conf *ServerConf
}

func NewTcpClient(conf *ServerConf)(cl *TcpClient,err error){
	cl = new(TcpClient)
	cl.conf = conf
	return
}

func (t *TcpClient)Recv(conn *net.TCPConn){
	defer func(){
		fmt.Println("断开服务器连接：")
	}()
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil{
			if err == io.EOF{
				return
			}
			return
		}
		if n > 0{
			fmt.Println("客户端收到数据：", string(buf))
		}
	}
}

func (t *TcpClient)Send(conn *net.TCPConn){
	
	var data []byte
	var startTime = time.Now().UnixNano()/1e6
	var endTime = time.Now().UnixNano()/1e6
	var cunt =0
	defer conn.Close()
	for{
		data = []byte(fmt.Sprintf("login game for now %d", time.Now().UnixNano()))
		// fmt.Println("发送一条数据：", string(data))
		conn.Write(data)
		// time.Sleep(time.Millisecond * 1)
		cunt++
		endTime = time.Now().UnixNano()/1e6
		
		if endTime-startTime >= 1000{
			fmt.Println("发送总量：", cunt)
			startTime = endTime
			cunt = 0
		}
	}
}

func (t *TcpClient)Start(){
	tcpadd, err := net.ResolveTCPAddr("tcp", t.conf.Host)
	// c, err := net.Dial("tcp", t.conf.Host)
	c, err := net.DialTCP("tcp", nil, tcpadd)
	if err != nil{
		fmt.Println("connect server failed", err)
	}
	go t.Recv(c)
	go t.Send(c)
}



func main(){
	conf := &ServerConf{
		Host:"192.168.8.154:20563",
	}

	for i:=0 ; i < 1; i++{
		go func(){
			client , _ := NewTcpClient(conf)
			client.Start()
		}()
		time.Sleep(time.Microsecond * 100)
	}
	
	select{}
}