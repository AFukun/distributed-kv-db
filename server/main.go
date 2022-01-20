package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/AFukun/distributed-kv-db/pb"
	"github.com/AFukun/distributed-kv-db/server/db"
	"google.golang.org/grpc"
)

type server struct{}

var database db.Database

func (s *server) Query(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	method := request.GetMethod()
	key := request.GetKey()

	if method == "GET" {
		value := database.Get(key)
		if value == 0 {
			fmt.Printf("Get key \"%s\" with null", key)
		} else {
			fmt.Printf("Get key \"%s\" with value %d\n", key, value)
		}
		return &pb.Response{Status: "VALUE", Value: int64(value)}, nil
	}

	if method == "PUT" {
		value := request.GetValue()
		database.Put(key, int(value))
		fmt.Printf("Put key \"%s\" to value %d\n", key, value)
		return &pb.Response{Status: "OK", Value: int64(0)}, nil
	}

	if method == "DELETE" {
		database.Delete(key)
		fmt.Printf("Delete key \"%s\"\n", key)
		return &pb.Response{Status: "OK", Value: int64(0)}, nil
	}

	return &pb.Response{Status: "ERROR", Value: int64(0)}, nil
}

func main() {
	port := os.Args[1]
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	database.Init()

	s := grpc.NewServer()
	pb.RegisterDatabaseServiceServer(s, &server{})

	fmt.Println("Server listening on Port " + port)
	if e := s.Serve(listener); e != nil {
		log.Fatal(e)
	}
}
