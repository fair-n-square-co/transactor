package db

import (
	"context"
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateUserFields struct {
	Email     string
	Username  string
	FirstName string
	LastName  string
}

type GetUserInput struct {
	Username string
}

type UserResponse struct {
	ID        uuid.UUID
	Email     string
	Username  string
	FirstName string
	LastName  string
}

type User struct {
	db *gorm.DB
}

func (u *User) CreateUser(ctx context.Context, user CreateUserFields) (uuid.UUID, error) {
	userModel := models.User{
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	result := u.db.Create(&userModel)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %v", result.Error)
	}
	return userModel.ID, nil
}

func (u *User) GetUser(ctx context.Context, in GetUserInput) (*UserResponse, error) {
	var user models.User
	result := u.db.Where("Username = ?", in.Username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user: %v", result.Error)
	}

	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil

}

func newUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}
