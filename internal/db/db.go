package db

import (
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/config"
	gormconfig "github.com/fair-n-square-co/transactions/internal/db/models/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	groupDBClient *Group
	userDBClient  *User
}

func (c *Client) GroupClient() *Group {
	return c.groupDBClient
}

func (c *Client) UserClient() *User {
	return c.userDBClient
}

func NewDB(cfg config.DatabaseConfig) (*Client, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), gormconfig.GetGormConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	group := newGroup(db)
	user := newUser(db)

	return &Client{
		group,
		user,
	}, nil
}
