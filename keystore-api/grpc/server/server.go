package main

import (
	"cloud-native-go/keystore-api/grpc/pb"
	"context"
)

type keyValueServer struct {
	pb.UnimplementedKeyValueServer
}

func (s *keyValueServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	value, err := store.GetValue(r.Key)
	return &pb.GetResponse{Value: value}, err

}

func (s *keyValueServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	err := store.PutValue(r.Key, r.Value)
	return &pb.PutResponse{}, err
}

func (s *keyValueServer) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := store.DeleteValue(r.Key)
	return &pb.DeleteResponse{}, err
}
