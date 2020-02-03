package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	proto "code_study_microservice/grpc-server-client/proto"
)

const port = 9000

type echoServer struct {}

func (e *echoServer) Hello(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	response := new(proto.Response)
	response.Msg = fmt.Sprintf("Hello %s", req.Name)

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	proto.RegisterEchoServer(server, &echoServer{})

	log.Printf("Server starting on port %v\n", port)
	server.Serve(listener)
}
