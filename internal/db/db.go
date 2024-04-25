package db

import (
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/db/models/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:generate mockgen -source=db.go -destination=mocks/mock_db.go -package=dbmocks

type Client interface {
	Group
}

type client struct {
	Group
}

func NewDB() (Client, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/transactions?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), config.GetGormConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	group := newGroup(db)

	return &client{
		group,
	}, nil
}
