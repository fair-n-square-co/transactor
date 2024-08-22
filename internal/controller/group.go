package controller

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/service/user/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/db"
	"github.com/google/uuid"
)

//go:generate mockgen -source=group.go -destination=mocks/mock_group.go -package=controllermocks

type GroupDBClient interface {
	// CreateGroup creates a new group in the database
	CreateGroup(ctx context.Context, group db.CreateGroupOptions) (uuid.UUID, error)
	ListGroups(ctx context.Context) (*db.GroupList, error)
}

// groupController is responsible for group requests.
type GroupController struct {
	dbClient GroupDBClient
}

// CreateGroup creates the group using the db package,
// and returns the proto response.
func (g *GroupController) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	options := db.CreateGroupOptions{
		Name: req.Name,
	}
	groupId, err := g.dbClient.CreateGroup(ctx, options)
	if err != nil {
		return nil, err
	}

	return &pb.CreateGroupResponse{
		GroupId: groupId.String(),
	}, nil
}

// ListGroups lists the groups using the db package,
// and returns the proto response.
func (g *GroupController) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	groups, err := g.dbClient.ListGroups(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListGroupsResponse{}
	for _, group := range groups.Groups {
		resp.Groups = append(resp.Groups, &pb.Group{
			Id:   group.ID.String(),
			Name: group.Name,
		})
	}

	return resp, nil
}

// NewGroupController creates a new instance of GroupController.
func NewGroupController(dbClient GroupDBClient) *GroupController {
	return &GroupController{
		dbClient: dbClient,
	}
}
