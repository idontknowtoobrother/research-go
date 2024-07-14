package main

import (
	"log"
	"net"
	"protobuf-practice/user"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcChildServerUser := &user.GrpcServerUser{
		Users:    make(map[string]*user.User),
		NextUuid: "",
	}
	user.RegisterUserManagerServer(grpcServer, grpcChildServerUser)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
