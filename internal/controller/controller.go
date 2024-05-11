package controller

import (
	"github.com/fair-n-square-co/transactions/internal/config"
	"github.com/fair-n-square-co/transactions/internal/db"
)

type DBClient interface {
	UserDBClient
	GroupDBClient
}

type Controller struct {
	u *UserController
	g *GroupController
}

func (c *Controller) UserController() *UserController {
	return c.u
}

func (c *Controller) GroupController() *GroupController {
	return c.g
}

func NewController(cfg config.DatabaseConfig) (*Controller, error) {
	dbClient, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Controller{
		u: NewUserController(dbClient.UserClient()),
		g: NewGroupController(dbClient.GroupClient()),
	}, nil
}
