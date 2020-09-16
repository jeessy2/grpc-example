package client

import (
	"context"
	"fmt"
	"grpc-example/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
)

func Helloworld() {
	firstTime := time.Now().Unix()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	for i := 0; i < 100000; i++ {
		r, err := c.SayHello(ctx, &proto.HelloRequest{Name: fmt.Sprint(i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
	log.Println("消耗时间(s):",time.Now().Unix()-firstTime)

}
