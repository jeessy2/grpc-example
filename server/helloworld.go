package server

import (
	"context"
	"grpc-example/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

const (
	port = ":9999"
)

// sayHello implements helloworld.GreeterServer.SayHello
func (ser *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Helloworld() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("listen on ", port)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
