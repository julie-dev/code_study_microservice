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
)

const (
	port = 9000

	crt        = "../cert/client.crt"
	key        = "../cert/client.key"
	rootCA    = "../cert/rootCA.crt"
	ServerName = "localhost"
)

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
		log.Fatal("Failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   ServerName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

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
