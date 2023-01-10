package main

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/theoriz0/hello-grpcgw/pkg/util"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	helloworldpb "github.com/theoriz0/hello-grpcgw/proto/helloworld"
)

var (
	CertPemPath = "certs/server.pem"
	CertKeyPath = "certs/server.key"
	Addr        = ":8090"
)

type helloService struct {
	helloworldpb.UnimplementedGreeterServer
}

func (s helloService) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create TLS config
	tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)
	var opts []grpc.ServerOption

	// Create a gRPC server object
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}
	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(grpcServer, &helloService{})

	//gw server
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, "dev.io")
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// Register Greeter
	err = helloworldpb.RegisterGreeterHandlerFromEndpoint(context.Background(), gwmux, ":8080", dopts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Serve gateway
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	server := &http.Server{
		Addr:      Addr,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}

	server.Serve(tls.NewListener(lis, tlsConfig))
}
