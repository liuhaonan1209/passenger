package main

import (
	"carRider-srv/internal/service"
	"flag"
	"log"
	"net"

	carrpc "carRider-srv/proto/rpc"

	_ "carRider-srv/init"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	carrpc.RegisterCartServer(s, &service.CartServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
