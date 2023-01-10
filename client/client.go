package main

import (
	helloworldpb "github.com/theoriz0/hello-grpcgw/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("certs/server.pem", "dev.io")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := helloworldpb.NewGreeterClient(conn)
	context := context.Background()
	body := &helloworldpb.HelloRequest{
		Name: "Grpc",
	}

	r, err := c.SayHello(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
