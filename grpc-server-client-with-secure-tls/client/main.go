package main

import (
	proto "code_study_microservice/grpc-server-client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	port = 9000

	certfile = "../cert/server.crt"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(certfile, "")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}
	defer conn.Close()

	client := proto.NewEchoClient(conn)
	response, err := client.Hello(context.Background(), &proto.Request{Name: "julie"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Msg)
}
