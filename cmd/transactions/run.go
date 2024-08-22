package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	usergw "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/service/user/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/cmd/transactions"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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

	usergw.RegisterGroupServiceServer(grpcServer, serverGroup.GroupServer)
	usergw.RegisterUserServiceServer(grpcServer, serverGroup.UserServer)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = usergw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, port, opts)
	if err != nil {
		return err
	}

	err = usergw.RegisterGroupServiceHandlerFromEndpoint(ctx, mux, port, opts)
	if err != nil {
		return err
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("starting http server on port %s", ":8081")
		return http.ListenAndServe(":8081", mux)
	})

	g.Go(func() error {
		log.Printf("listening on port %s", port)
		return grpcServer.Serve(lis)
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("server error: %v", err)
		return err
	}

	return nil
}
