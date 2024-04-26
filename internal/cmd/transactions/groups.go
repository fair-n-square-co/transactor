package transactions

import (
	"context"
	"fmt"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/config"
	"github.com/fair-n-square-co/transactions/internal/controller"
	"github.com/fair-n-square-co/transactions/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupsServer struct {
	pb.UnimplementedGroupServiceServer
	controller controller.Controller
	config     config.Config
}

func (g *GroupsServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}
	res, err := g.controller.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GroupsServer) ListGroups(ctx context.Context, request *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	res, err := g.controller.ListGroups(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewGroupsServer() (pb.GroupServiceServer, error) {
	config := config.NewConfig()
	dbClient, err := db.NewDB(config.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %v", err)
	}
	controller, err := controller.NewController(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %v", err)
	}
	return &GroupsServer{
		controller: controller,
		config:     config,
	}, nil
}
