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

func (s *Server) AddBlock(ctz context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, err) {
	block := s.Blockchain.AddBlock(in.Data)

	return &proto.AddBlockResponse{
		Hash: block.Hash,

	}, nil

}

func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, err) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			b.PrevBlockHash: b.PrevBlockHash
			Hash: b.Hash,
			Data: b.Hash,
		})
	}

	return resp, nil
}
