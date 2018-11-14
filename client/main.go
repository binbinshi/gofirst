package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gofirstdemo/proto"
)

const Address  = "127.0.0.1:9898"


func main() {
	conn, err :=grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer  conn.Close()

	c := pb.NewBinBinClient(conn)

	reqBody := new(pb.BinBinRequest)
	reqBody.Name = "grpc-bin"
	r, err := c.SayHi(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)
	}
