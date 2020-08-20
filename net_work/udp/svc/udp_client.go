package svc

import (
	"fmt"
	"net"
)

type UdpClient struct {
	host    string
	port    string
	laAddr  *net.UDPAddr
	raAddr  *net.UDPAddr
	udpConn *net.UDPConn
}

func NewUdpClient(host string, port string) *UdpClient {
	us := &UdpClient{
		host: host,
		port: port,
	}

	raAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}
	us.raAddr = raAddr
	return us
}

func (sel *UdpClient) SetLocalAddr(addr string) error {
	laAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	sel.laAddr = laAddr
	return nil
}

func (sel *UdpClient) Connect() {
	laAddr := sel.laAddr
	raAddr := sel.raAddr
	conn, err := net.DialUDP("udp", laAddr, raAddr)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("链接服务成功：", raAddr.String())
	sel.udpConn = conn
}

func (sel *UdpClient) Send() error {
	if sel.udpConn == nil {
		sel.Connect()
	}
	conn := sel.udpConn
	_, err := conn.Write([]byte("中彩票"))
	return err
}

func (sel *UdpClient) Recv() {
	updConn := sel.udpConn
	data := make([]byte, 1024)
	n, err := updConn.Read(data)
	if err != nil {
		fmt.Println("接收消息失败：", err)
		return
	}

	data = data[:n]
	//fmt.Println("接收到来自", raddr.String(), "的消息:", string(data))
	fmt.Println("接收到来自", "的消息:", string(data))

}
