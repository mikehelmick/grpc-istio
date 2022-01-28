package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	cpb "github.com/mikehelmick/grpc-istio/pkg/counter/pb"
	"github.com/mikehelmick/grpc-istio/pkg/counter/server"

	grpc "google.golang.org/grpc"
)

var port = flag.Int("port", 3232, "port number to listen on")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	cpb.RegisterEchoServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
