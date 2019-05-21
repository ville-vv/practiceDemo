package sserver

import (
	"context"
	"fmt"
	"net"
	pb "practiceDemo/grcp_gateway_work/simple"

	"google.golang.org/grpc"
)

type ServerSimple struct {
	addr  string
	grpcS *grpc.Server
}

func NewServerSimple(addr string) *ServerSimple {
	ss := new(ServerSimple)
	ss.grpcS = grpc.NewServer()
	ss.addr = addr
	return ss
}

func (s *ServerSimple) Start() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Println("server start failed ：", err)
		return
	}
	fmt.Println("server run ", s.addr)
	pb.RegisterSimpleServerServer(s.grpcS, s)
	s.grpcS.Serve(lis)
}

// 接口继承 protobuf的grpc 中的 SimpleServerServer 接口。 请求数据入口是在这里
func (s *ServerSimple) Stream(ctx context.Context, r *pb.Package_Request) (w *pb.Package_Response, err error) {
	fmt.Println("请求参数：", r)
	w = new(pb.Package_Response)
	w.Code = 200
	w.Message = "ok"
	w.Data = []byte("接口继承 protobuf的grpc 中的 SimpleServerServer 接口。 请求数据入口是在这里")
	return w, nil
}
