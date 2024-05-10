package main

import (
	"log"
	"net"
	"os"

	v1alpha1 "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/cmd/transactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func run() {
	// get PORT from env variables
	// if not set, use default port
	port := ":8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = ":" + envPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("unable to listen on port ", port)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Println(err)
		}
	}(lis)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register API v1
	serverGroup, err := transactions.NewServerGroup()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	v1alpha1.RegisterGroupServiceServer(grpcServer, serverGroup.GroupServer)
	v1alpha1.RegisterUserServiceServer(grpcServer, serverGroup.UserServer)

	log.Printf("listening on port %s", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed at: %v", err)
	}
}
