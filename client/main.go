package main

import "flag"

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listflag := flag.Bool("lis", false, "get new blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fataf("canot dial server: %v", err)
	}

	client := proto

}
