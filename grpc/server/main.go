package main

import (
	"cloud-native-go/common"
	"cloud-native-go/grpc/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var store = common.Store{
	Data: make(map[string]string),
}

func main() {
	port := 8000
	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &keyValueServer{})

	lis, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatalf("failed to listen: %s\n", err)
	}

	log.Printf("starting gRPC server on port %d\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
