package main

import (
	"fmt"
	pb "gofirstdemo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	Address = "127.0.0.1:9898"
)

type binbinService struct{}

func (b binbinService) SayHi(ctx context.Context, in *pb.BinBinRequest) (*pb.BinBinReply, error) {
	resp := new(pb.BinBinReply)
	resp.Message = "hello " + in.Name + "."
	return resp, nil
}

var BinBinServer = binbinService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("failed to listen:%v", err)
	}
	//实现grpc server
	s := grpc.NewServer()
	//注册binbinServer为客户端提供服务
	pb.RegisterBinBinServer(s, BinBinServer)
	fmt.Println("Listen on" + Address)
	s.Serve(listen)
}
