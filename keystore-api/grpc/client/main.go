package main

import (
	"cloud-native-go/keystore-api/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
)

func main() {
	// establish connection to gRPC server
	conn, err := grpc.Dial(
		"localhost:8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %s\n", err)
	}
	defer conn.Close()

	// created new client for connection
	client := pb.NewKeyValueClient(conn)

	// set inputs from flags
	var action, key, value string
	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], " ")
	}

	// perform specified action
	switch action {
	case "get":
		r, err := client.Get(context.Background(), &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("could not get value key %s: %s\n", key, err)
		}
		fmt.Printf("%s: %s\n", key, r.Value)

	case "put":
		_, err := client.Put(context.Background(), &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("could not put key %s: %s\n", key, err)
		}
		fmt.Printf("put %s\n", key)

	case "delete":
		_, err := client.Delete(context.Background(), &pb.DeleteRequest{Key: key})
		if err != nil {
			log.Fatalf("could not delete key %s: %v\n", key, err)
		}
		fmt.Printf("deleted %s\n", key)

	default:
		fmt.Println("usage: go run [get|put]delete] KEY [VALUE]")
	}
}
