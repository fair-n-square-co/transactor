package transactions

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
)

type groupsServer struct {
	pb.UnimplementedGroupServiceServer
}

func (g *groupsServer) CreateGroup(context.Context, *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	/* TODO
		1. validate data
		2. create group in db
		3. return success
		4. Failure scenarios
	*/ 
	
	return &pb.CreateGroupResponse{}, nil
}


func (g *groupsServer) ListGroups(context.Context, *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	
	return &pb.ListGroupsResponse{}, nil
}
