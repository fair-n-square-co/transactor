package controller

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/service/user/v1alpha1"
	usertypepb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/type/user/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/db"
	"github.com/google/uuid"
)

type UserDBClient interface {
	CreateUser(ctx context.Context, user db.CreateUserFields) (uuid.UUID, error)
	GetUser(ctx context.Context, in db.GetUserOptions) (*db.UserResponse, error)
}

type UserController struct {
	dbClient UserDBClient
}

func (u *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fields := db.CreateUserFields{
		Email:     req.Email,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	userId, err := u.dbClient.CreateUser(ctx, fields)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		UserId: userId.String(),
	}, nil
}

func (u *UserController) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	input := db.GetUserOptions{
		Username: in.Username,
	}

	user, err := u.dbClient.GetUser(ctx, input)
	if err != nil {
		return nil, err
	}

	var userResponse = &usertypepb.User{
		UserId:    user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return &pb.GetUserResponse{
		User: userResponse,
	}, nil
}

// NewUserController creates a new instance of UserController.
func NewUserController(dbClient UserDBClient) *UserController {
	return &UserController{
		dbClient: dbClient,
	}
}
