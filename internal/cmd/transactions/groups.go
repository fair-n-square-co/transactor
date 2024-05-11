package transactions

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source=group.go -destination=mocks/mock_group.go -package=controllermocks

type GroupController interface {
	CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error)
	ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error)
}

type GroupServer struct {
	pb.UnimplementedGroupServiceServer
	controller GroupController
	config     config.Config
}

func (g *GroupServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}
	res, err := g.controller.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GroupServer) ListGroups(ctx context.Context, request *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	res, err := g.controller.ListGroups(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewGroupServer(groupController GroupController) (*GroupServer, error) {
	return &GroupServer{
		controller: groupController,
	}, nil
}
