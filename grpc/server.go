package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"study/grpc/proto"
)

const PORT = "9001"

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: r.GetRequest() + "Service"}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen() err: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve() err: %v", err)
	}
}
