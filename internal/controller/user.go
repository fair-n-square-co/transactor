package controller

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/db"
)

type UserController interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error)
}

type userController struct {
	dbClient db.Client
}

func (u *userController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fields := db.CreateUserFields{
		Email: req.Email,
		Username: req.Username,
		FirstName: req.FirstName,
		LastName: req.LastName,
	}

	userId, err := u.dbClient.CreateUser(ctx, fields)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		UserId: userId.String(),
	}, nil
}

func (u *userController) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	input := db.GetUserInput{
		Username: in.Username,
	}

	user, err := u.dbClient.GetUser(ctx, input)
	if err != nil {
		return nil, err
	}

	var userResponse = &pb.User{
		UserId:    user.ID.String(),
		Username:  user.Username,
		Email:	   user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return &pb.GetUserResponse{
		User: userResponse,
	}, nil
}

// NewUserController creates a new instance of UserController.
func NewUserController(dbClient db.Client) UserController {
	return &userController{
		dbClient: dbClient,
	}
}