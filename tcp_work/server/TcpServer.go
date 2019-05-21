package main

import (
	"runtime"
	"time"
	"fmt"
	"net"
	"sync"
	"github.com/soheilhy/cmux"
	"net/http"
	"golang.org/x/net/websocket"
	"io"
)

type PackageData struct{
	Id uint64
	Body []byte
}


type ServerConf struct{
	Host string
}

type ConnChan struct{
	id uint64
	conn net.Conn
	stopCh chan int
	isStop bool
	lock sync.Mutex
}
func (c *ConnChan)IsStop()bool{
	return c.isStop
}

func(c *ConnChan)Recv()([]byte, error){
	var(
		data = make([]byte, 1*1024)
	)
	n, err := c.conn.Read(data)
	if err != nil{
		fmt.Println("disconnect：", err)
		c.isStop = true
		if err == io.EOF{
			return data,nil
		}
		return nil, err
	}
	if(n > 0){
		data = data[:n]
	}
	return data, nil
}

func (c *ConnChan)Send(buf []byte)(err error){
	_, err = c.conn.Write(buf)
	if err != nil{
		fmt.Println("发送数据失败：", err)
	}
	return 
}


type TcpServer struct{
	idCnt	uint64
	listener net.Listener
	sconf	*ServerConf
	connMap	map[uint64]*ConnChan
	recvCh chan PackageData
	sendCh chan PackageData
	lock sync.Mutex
}

func NewTcpServer(config *ServerConf)(t *TcpServer, err error){
	t = new(TcpServer)
	t.connMap = make(map[uint64]*ConnChan)
	t.sconf = config
	t.recvCh = make(chan PackageData)
	t.sendCh = make(chan PackageData)
	return
}

func EchoServer(ws *websocket.Conn) {
    if _, err := io.Copy(ws, ws); err != nil {
        panic(err)
    }
}

func (t *TcpServer)serveWS(l net.Listener) {
    s := &http.Server{
        Handler: websocket.Handler(EchoServer),
    }
    if err := s.Serve(l); err != cmux.ErrListenerClosed {
        panic(err)
    }
}

func (t *TcpServer)Start()(err error){
	lis, err := net.Listen("tcp", t.sconf.Host)
	if err != nil{
		fmt.Println("tcp server start failed", err)
		return
	}
	fmt.Println("Server run success")
	m := cmux.New(lis)
	tcpl := m.Match(cmux.Any())
	t.SetListener(tcpl)
	go t.ToServer(tcpl)

	m.Serve()
	return
}

func (t *TcpServer)ToServer(l net.Listener){
	defer l.Close()

	for{
		conn , err := l.Accept()
		if err != nil{
			fmt.Println("server connection error:", err)
			continue
		}
		fmt.Println("收到请求连接：",conn.RemoteAddr())

		go t.NoChanHandler(conn)

		// go t.ServerConnHandler(conn)
	}
}

func (t *TcpServer)ServerConnHandler(conn net.Conn){
	c := &ConnChan{
		conn:conn,
		isStop:false,
		stopCh:make(chan int),
	}

	t.setConnMap(0, c)

	defer func(){
		t.delConnMap(c.id)
	}()

	go t.ReceiveMs(c)
	go t.SendMs(c)

	for{
		if c.IsStop(){
			return
		}
		select{
		case p :=<- t.recvCh:
			fmt.Println("收到数据：", string(p.Body))
			rsp := PackageData{}
			rsp.Id = 0
			rsp.Body = []byte(fmt.Sprintf("Server Response %d", time.Now().UnixNano()))
			// c.Send(rsp.Body)
			t.sendCh <- rsp
		case <- time.After(time.Second*1):
			break
		}
		// time.Sleep(time.Second*3)
	}
}

func (t *TcpServer)NoChanHandler(conn net.Conn){
	c := &ConnChan{
		conn:conn,
		isStop:false,
		stopCh:make(chan int),
	}

	t.setConnMap(0, c)
	var cout = 0
	defer func(){
		fmt.Println("收到客户端数据总数：", cout)
		t.delConnMap(c.id)
	}()

	
	var startTime = time.Now().Unix()
	var endTime = time.Now().Unix()
	for{
		data := make([]byte, 128*1024)
		_, err :=conn.Read(data)
		if err != nil{
			fmt.Println("disconnect：", err);
			if err == io.EOF{
				return 
			}
			return 
		}
		endTime = time.Now().Unix()
		// fmt.Println("收到客户端数据：", string(data[:n]))
		cout++
		if endTime-startTime >= 1{
			fmt.Println("收到客户端数据总数：", cout)
			startTime = endTime
			cout = 0
		}
			// conn.Write([]byte(fmt.Sprintf("Server Response %d", time.Now().UnixNano())))
	}
}


func (t *TcpServer)ReceiveMs(conn *ConnChan)(err error){
	for{
		data , err := conn.Recv()
		
		if err != nil{
			break
		}
		// fmt.Println("收到消息：",string(data))
		pkg := PackageData{
			Id : conn.id,
			Body: data[:],
		}
		t.recvCh <- pkg
		// conn.Send([]byte(fmt.Sprintf("Server Response %d", time.Now().UnixNano())))
	}
	return
}

func (t *TcpServer)SendMs(conn *ConnChan)(err error){
	for{
		if conn.IsStop() {
			return
		}
		select{
		case p :=<- t.sendCh:
			conn.Send(p.Body)
		case <- time.After(time.Second*1):
			break
		}
	}
}

func (t *TcpServer)SetListener(listener net.Listener){
	t.listener = listener
}

func (t *TcpServer)setConnMap(key uint64, conn *ConnChan){
	t.lock.Lock()
	t.idCnt++
	key = t.idCnt
	conn.id = key
	t.connMap[key] = conn
	fmt.Println("添加连接Map:", key, len(t.connMap))
	t.lock.Unlock()
}

func (t *TcpServer)delConnMap(key uint64){
	t.lock.Lock()
	delete (t.connMap, key)
	fmt.Println("离开：", key, len(t.connMap))
	t.lock.Unlock()
	return
}


func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf := &ServerConf{
		Host:"192.168.8.154:20563",
	}
	server, _ := NewTcpServer(conf)

	server.Start()
	fmt.Println("结束：")
}