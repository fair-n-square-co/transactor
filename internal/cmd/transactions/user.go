package transactions

import (
	"context"
	"log"
	"strings"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error)
}

type UserServer struct {
	pb.UnimplementedUserServiceServer
	controller UserController
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

func (u *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if strings.TrimSpace(req.Username) == "" {
		return nil, status.Error(codes.InvalidArgument, "Username cannot be empty")
	}

	res, err := u.controller.GetUser(ctx, req)
	if err != nil {
		log.Printf("Get User error: %v", err)
		return nil, status.Error(codes.Internal, "Error occured while getting user information")
	}

	return res, nil
}

func NewUserServer(controller UserController) (*UserServer, error) {
	return &UserServer{
		controller: controller,
	}, nil
}
