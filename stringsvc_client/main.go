package main

import (
	"log"
	"os"

	pb "../stringsvc_contract"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultString = "TessTTT"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStringsvcClient(conn)

	// Contact the server and print out its response.
	inputnotformatted := defaultString
	if len(os.Args) > 1 {
		inputnotformatted = os.Args[1]
	}
	response, err := c.UpperCase(context.Background(), &pb.StringRequest{Messagenotformatted: inputnotformatted})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s transform in: %s", defaultString, response.Messageformatted)

	response, err = c.LowerCase(context.Background(), &pb.StringRequest{Messagenotformatted: inputnotformatted})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s transform in: %s", defaultString, response.Messageformatted)
}
