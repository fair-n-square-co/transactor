package controller

import "github.com/fair-n-square-co/transactions/internal/db"

type Controller interface {
	GroupController
}

type controller struct {
	GroupController
}

func NewController(dbClient db.Client) (Controller, error) {
	return &controller{
		NewGroupController(dbClient),
	}, nil
}
