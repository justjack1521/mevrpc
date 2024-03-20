package mevrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunGRPCServer(port string, register func(server *grpc.Server)) {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	RunGRPCServerOnAddr(addr, register)
}

func RunGRPCServerWithOptions(port string, register func(server *grpc.Server), options ...grpc.ServerOption) {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	RunGRPCServerOnAddrWithOptions(addr, register, options...)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {

	svr := grpc.NewServer()

	registerServer(svr)

	reflection.Register(svr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Now listening on: %s", listen.Addr()))

	err = svr.Serve(listen)
	if err != nil {
		panic(err)
	}

}

func RunGRPCServerOnAddrWithOptions(addr string, registerServer func(server *grpc.Server), options ...grpc.ServerOption) {
	svr := grpc.NewServer(options...)

	registerServer(svr)

	reflection.Register(svr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Now listening on: %s", listen.Addr()))

	err = svr.Serve(listen)
	if err != nil {
		panic(err)
	}
}
