package main

import (
	proto "code_study_microservice/grpc-server-client/proto"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

const (
	port = 9000

	crt = "../cert/server.crt"
	key = "../cert/server.key"
	rootCA  = "../cert/rootCA.crt"
)

type echoServer struct{}

func (e *echoServer) Hello(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	response := new(proto.Response)
	response.Msg = fmt.Sprintf("Hello %s", req.Name)

	fmt.Print(".")

	return response, nil
}

func main() {
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(rootCA)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("Failed to append certs") //?
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
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
