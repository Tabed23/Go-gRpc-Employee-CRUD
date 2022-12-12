package cmd

import (
	"employee-grpc/pb"
	"employee-grpc/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartApplication() {
	opts := []grpc.ServerOption{}
	rpcServer := grpc.NewServer(opts...)
	pb.RegisterEmployeeServiceServer(rpcServer, &service.Server{})
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		fmt.Printf("error with server %v", err)
	}
	if err = rpcServer.Serve(lis); err != nil {
		log.Fatalf("cannot server the request %v", err)
	}
}
