package main

import (
	"log"
	"net"
	"strings"

	pb "../stringsvc_contract"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement Stringcontract.StringServer.
type server struct{}

// UpperCase implements stringsvc_contract.UpperCase
func (s *server) UpperCase(ctx context.Context, in *pb.StringRequest) (*pb.StringReply, error) {
	return &pb.StringReply{Messageformatted: "UpperCase " + strings.ToUpper(in.Messagenotformatted)}, nil
}

// LowerCase implements stringsvc_contract.LowerCase
func (s *server) LowerCase(ctx context.Context, in *pb.StringRequest) (*pb.StringReply, error) {
	return &pb.StringReply{Messageformatted: "UpperCase " + strings.ToLower(in.Messagenotformatted)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringsvcServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
