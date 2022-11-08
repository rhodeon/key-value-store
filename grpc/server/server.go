package main

import (
	pb2 "cloud-native-go/grpc/pb"
	"context"
)

type keyValueServer struct {
	pb2.UnimplementedKeyValueServer
}

func (s *keyValueServer) Get(ctx context.Context, r *pb2.GetRequest) (*pb2.GetResponse, error) {
	value, err := store.GetValue(r.Key)
	return &pb2.GetResponse{Value: value}, err

}

func (s *keyValueServer) Put(ctx context.Context, r *pb2.PutRequest) (*pb2.PutResponse, error) {
	err := store.PutValue(r.Key, r.Value)
	return &pb2.PutResponse{}, err
}

func (s *keyValueServer) Delete(ctx context.Context, r *pb2.DeleteRequest) (*pb2.DeleteResponse, error) {
	err := store.DeleteValue(r.Key)
	return &pb2.DeleteResponse{}, err
}
