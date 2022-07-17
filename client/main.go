package main

import (
	"context"
	"flag"

	"github.com/0xlax/gRPC"
)


var client proto.BLockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listflag := flag.Bool("lis", false, "get new blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fataf("canot dial server: %v", err)
	}

	client := proto.NewBLockchainClient(conn)

	if *addFlag {
		AddBlock()
	}

	iff *listflag{
		getBlockchain()
	}

}


func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().ToString()
	})
	if err != nill {
		log.Fatalf("unable to add block")
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nill {
		log.Fatalf("unable to get")
	}
	log.Printf("blocks")
	// Print WIP- Print Current.block contents
}