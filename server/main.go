package main

import (
	"context"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != null {
		log.Fatalf("unable  listern on 8080 port: %v", error)
	}
	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)

}

type Server struct{}

func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, err) {
	return new(proto.AddBlockResponse), nil

}

func (s *Server) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, err) {
	return new(proto.GetBlockchainResponse), nil
}
