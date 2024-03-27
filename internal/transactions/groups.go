package transactions

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type groupsServer struct {
	pb.UnimplementedGroupServiceServer
}

func (g *groupsServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	/* TODO
	1. validate data
	2. create group in db
	3. return success
	4. Failure scenarios
	*/
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}

func (g *groupsServer) ListGroups(ctx context.Context, request *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}

func NewGroupsServer() pb.GroupServiceServer {
	return &groupsServer{}
}
