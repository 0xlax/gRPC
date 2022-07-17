package main

import (
	"context"
	"log"
	"net"

	"github.com/0xlax/gRPC/server/blockchain"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != null {
		log.Fatalf("unable  listern on 8080 port: %v", error)
	}
	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBLockchain(),
	})
	srv.Serve(listener)

}

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, err) {
	return new(proto.AddBlockResponse), nil

}

func (s *Server) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, err) {
	return new(proto.GetBlockchainResponse), nil
}
