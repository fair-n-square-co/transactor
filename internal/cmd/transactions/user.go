package transactions

import (
	"context"
	"fmt"
	"log"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/config"
	"github.com/fair-n-square-co/transactions/internal/controller"
	"github.com/fair-n-square-co/transactions/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	controller controller.UserController
	config config.Config
}

func (u *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.Email == "" || req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "Email/Username cannot be empty")
	}

	res, err := u.controller.CreateUser(ctx, req)
	if err != nil {
		log.Printf("Create User error: %v", err)
		return nil, status.Error(codes.Internal, "Error occured while creating user")
	}

	return res, nil
}

func NewUserServer() (pb.UserServiceServer, error) {
	config := config.NewConfig()
	dbClient, err := db.NewDB(config.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %v", err)
	}
	controller, err := controller.NewController(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %v", err)
	}
	return &UserServer{
		controller: controller,
		config:     config,
	}, nil
}

