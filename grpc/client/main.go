package main

import (
	"context"
	"grpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	req := &proto.Request{A: int64(6), B: int64(5)}
	resp, err := client.Add(context.TODO(), req)
	log.Println(resp.GetResult())
}
