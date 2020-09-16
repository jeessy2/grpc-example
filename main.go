package main

import (
	"grpc-example/client"
	"grpc-example/server"
	"os"
)

func main() {

	switch os.Args[1] {
	case "server":
		server.Helloworld()
	default:
		client.Helloworld()
	}

}
