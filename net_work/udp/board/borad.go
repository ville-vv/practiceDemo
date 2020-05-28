package board

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func BoardCast1() {

	//f, err := os.Open("./image_01.jpg")
	//if err != nil {
	//	fmt.Println("打开文件失败", err)
	//	return
	//}
	//fData, err := ioutil.ReadAll(f)
	//if err != nil {
	//	fmt.Println("读取文件失败", err)
	//	return
	//}
	//f.Close()
	//fmt.Println("读取到文件数据长度：", len(fData))
	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	laddr := net.UDPAddr{
		IP:   net.IPv4(172, 16, 5, 48),
		Port: 3000,
	}
	// 这里设置接收者的IP地址为广播地址
	// 广播地址可以使 ifconfig 或者 ipconfig 查看当前网络接口所在的网段对应的 广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(172, 16, 5, 255),
		Port: 30000,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()
	stop := make(chan int)
	for i := 0; i < 1; i++ {
		go func(c *net.UDPConn, int2 int) {
			for {
				select {
				case <-stop:
					return
				default:
				}
				data := fmt.Sprintf("%d当前时间为：%s", int2, time.Now().String())
				fmt.Println("发送数据：", data)
				if _, err := conn.Write([]byte(data)); err != nil {
					fmt.Println("发送广播数据失败：", err)

				}
				//if _, err := conn.Write(fData[:48920]); err != nil {
				//	fmt.Println("发送文件数据失败：", err)
				//}
				time.Sleep(time.Second * 3)
			}
		}(conn, i)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	close(stop)
	time.Sleep(time.Second * 1)

}

func RecBoardCast() {
	// 本地监听的ip信息
	laddr := &net.UDPAddr{
		IP:   net.IPv4(172, 16, 5, 255),
		Port: 30000,
	}
	updConn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("监听地址失败：", laddr.String())
		return
	}

	for {
		data := make([]byte, 1024)
		n, raddr, err := updConn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("接收消息失败：", err)
			return
		}
		data = data[:n]
		fmt.Println("接收到来自", raddr.String(), "的消息:", string(data))
	}
}
