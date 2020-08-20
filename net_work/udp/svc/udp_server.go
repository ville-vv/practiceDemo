package svc

import (
	"fmt"
	"net"
)

type UdpServer struct {
	host    string
	port    string
	laAddr  *net.UDPAddr
	raAddr  []*net.UDPAddr
	udpConn *net.UDPConn
}

func NewUdpServer(host string, port string) (*UdpServer, error) {
	ur := &UdpServer{
		host: host,
		port: port,
	}
	laAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}
	ur.laAddr = laAddr

	return ur, ur.Listen()
}

func (u *UdpServer) Listen() error {
	laAddr := u.laAddr
	updConn, err := net.ListenUDP("udp", laAddr)
	if err != nil {
		fmt.Println("监听地址失败：", laAddr.String())
		return err
	}
	fmt.Println("监听地址成功：", laAddr.String())
	u.udpConn = updConn
	return nil
}

func (u *UdpServer) AddRemoteAddr(addr string) error {
	raAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	u.raAddr = append(u.raAddr, raAddr)
	return nil
}

func (u *UdpServer) ReadFromUdp() {
	updConn := u.udpConn
	data := make([]byte, 1024)
	//n, raddr, err := updConn.ReadFromUDP(data)
	//if err != nil {
	//	fmt.Println("接收消息失败：", err)
	//	return
	//}
	n, raddr, err := updConn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收消息失败：", err)
		return
	}

	data = data[:n]
	//fmt.Println("接收到来自", raddr.String(), "的消息:", string(data))
	fmt.Println("接收到来自", raddr.String(), "的消息:", string(data))
	_ = u.addClient(raddr)
}

func (u *UdpServer) addClient(addr *net.UDPAddr) error {
	raAddr := &net.UDPAddr{
		IP:   addr.IP,
		Port: 20000,
	}
	u.raAddr = append(u.raAddr, raAddr)
	return nil
}

func (u *UdpServer) Send() error {
	udpConn := u.udpConn
	for _, raAddr := range u.raAddr {
		if _, err := udpConn.WriteToUDP([]byte("收到消息，反馈"), raAddr); err != nil {
			fmt.Println("写入消息错误", err)
			return err
		}
	}
	return nil
}
