package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "E-Learning/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	menu()
}

//ConnectServer untuk connect ke proto
func ConnectServer() pb.ConnectClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	return pb.NewConnectClient(conn)
}

func menu() {
	var input int
	fmt.Println("ELEARNING PROTOTYPE")
	fmt.Println("1. Coba panggil server")
	fmt.Println("Pilih menu: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		panggilServer()
	default:
		fmt.Println("Input menu salah, mohon input angka yang sesuai")
	}
}

func panggilServer() {
	user := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	msg, err := user.ConnectToServer(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Message", msg.GetMessage())
	menu()
}
