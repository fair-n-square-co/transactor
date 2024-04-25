package controller

import "github.com/fair-n-square-co/transactions/internal/db"

//go:generate mockgen -source=controller.go -destination=mocks/mock_controller.go -package=controllermocks

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
