package main

import (
	"context"
	"fmt"
	"net/http"
	gw "practiceDemo/grcp_gateway_work/simple"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// logicAddr grpc 服务地址
// httpAddr http 服务地址
func run(logicAddr, httpAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := logicAddr

	err := gw.RegisterSimpleServerHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		fmt.Println("RegisterSimpleServerHandlerFromEndpoint error:", err)
		return err
	}
	fmt.Println("http 服务启动：", httpAddr)
	err = http.ListenAndServe(httpAddr, mux)
	if err != nil {
		fmt.Println("ListenAndServe failed ", err)
		return err
	}
	return nil
}

func main() {

	if err := run("0.0.0.0:32111", "0.0.0.0:32112"); err != nil {
		fmt.Println("run error", err)
		return
	}
}
