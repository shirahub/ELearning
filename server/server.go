package main

import (
	pb "E-Learning/proto"
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedConnectServer
}

func (s *server) ConnectToServer(ctx context.Context, void *empty.Empty) (*pb.ServerReply, error) {
	return &pb.ServerReply{Message: "Client is connected to Server"}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterConnectServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
