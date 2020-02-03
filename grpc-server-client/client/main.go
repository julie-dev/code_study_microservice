package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"

	proto "code_study_microservice/grpc-server-client/proto"
)

const port = 9000

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	client := proto.NewEchoClient(conn)
	response, _ := client.Hello(context.Background(), &proto.Request{Name:"julie"})
	if err != nil {
		log.Fatal("Error calling service: ", err)
		os.Exit(1)
	}

	if response == nil {
		//Exception of localhost case
		log.Fatal("No data available")
		os.Exit(1)
	}

	fmt.Println(response.Msg)
}