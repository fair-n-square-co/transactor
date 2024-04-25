package controller

import (
	"context"
	"testing"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/db"
	dbmocks "github.com/fair-n-square-co/transactions/internal/db/mocks"
	"github.com/google/uuid"
	tassert "github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupControllerCreateGroup(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockDBClient := dbmocks.NewMockClient(ctrl)
	groupController := NewGroupController(mockDBClient)

	ctx := context.Background()
	req := &pb.CreateGroupRequest{
		Name: "test",
	}

	mockDBClient.EXPECT().CreateGroup(ctx, db.CreateGroupOptions{
		Name: "test",
	}).Return(uuid.New(), nil)

	_, err := groupController.CreateGroup(ctx, req)
	assert.NoError(err)
}
