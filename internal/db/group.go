package db

import (
	"context"
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/fair-n-square-co/transactions/internal/db/models/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateGroupOptions struct {
	Name string
}

type GroupData struct {
	ID    uuid.UUID
	Name  string
	Users []UserResponse
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
	result := g.db.Model(&groups).Preload("Users").Find(&groups)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list groups: %v", result.Error)
	}
	for _, group := range groups {
		users := make([]UserResponse, 0, len(group.Users))
		for _, user := range group.Users {
			users = append(users, UserResponse{
				ID:        user.ID,
				Email:     user.Email,
				Username:  user.Username,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			})
		}
		groupList.Groups = append(groupList.Groups, GroupData{
			ID:    group.ID,
			Name:  group.Name,
			Users: users,
		})
	}
	return &groupList, nil
}

func (g *Group) UpdateUsersInGroup(ctx context.Context, groupID uuid.UUID, userIDs []uuid.UUID) error {
	// Get group from groupId
	var group models.Group
	result := g.db.First(&group, "id = ?", groupID)
	if result.Error != nil {
		return fmt.Errorf("failed to find group: %v", result.Error)
	}

	// Get users from userIDs
	users := make([]models.User, 0, len(userIDs))
	for _, id := range userIDs {
		users = append(users, models.User{PrimaryKey: base.PrimaryKey{ID: id}})
	}

	return g.db.Model(&group).Association("Users").Replace(users)
}

func newGroup(db *gorm.DB) *Group {
	return &Group{
		db: db,
	}
}
