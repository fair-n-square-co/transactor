package main

import (
	"log"
	"net"

	v1alpha1 "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/pkg/transactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func run() {
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
	server := grpc.NewServer()
	reflection.Register(server)

	// Register API v1
	service := transactions.NewTransactionsServer()
	v1alpha1.RegisterHelloServiceServer(server, service)

	log.Printf("listening on port %s", port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed at: %v", err)
	}
}
