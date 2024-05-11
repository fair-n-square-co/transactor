package db

import (
	"context"
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateGroupOptions struct {
	Name string
}

type GroupData struct {
	ID   uuid.UUID
	Name string
}

type GroupList struct {
	Groups []GroupData
}

type Group struct {
	db *gorm.DB
}

func (g *Group) CreateGroup(ctx context.Context, groupOptions CreateGroupOptions) (uuid.UUID, error) {
	groupModel := models.Group{
		Name: groupOptions.Name,
	}
	result := g.db.Create(&groupModel)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create group: %v", result.Error)
	}
	return groupModel.ID, nil
}

func (g *Group) ListGroups(ctx context.Context) (*GroupList, error) {
	var groups []models.Group
	var groupList GroupList
	result := g.db.Find(&groups)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list groups: %v", result.Error)
	}
	for _, group := range groups {
		groupList.Groups = append(groupList.Groups, GroupData{
			ID:   group.ID,
			Name: group.Name,
		})
	}
	return &groupList, nil
}

func newGroup(db *gorm.DB) *Group {
	return &Group{
		db: db,
	}
}
