package main

import (
	proto "code_study_microservice/grpc-server-client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	port = 9000

	certfile = "../cert/server.crt"
	keyfile = "../cert/server.key"
)

type echoServer struct{}

func (e *echoServer) Hello(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	response := new(proto.Response)
	response.Msg = fmt.Sprintf("Hello %s", req.Name)

	fmt.Print(".")

	return response, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile(certfile, keyfile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	server := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterEchoServer(server, &echoServer{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Printf("Server starting on port %v\n", port)
	if err = server.Serve(listener); err != nil {
		log.Fatal("Error: ", err)
	}
}
