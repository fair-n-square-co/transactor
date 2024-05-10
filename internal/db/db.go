package db

import (
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/config"
	gormconfig "github.com/fair-n-square-co/transactions/internal/db/models/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:generate mockgen -source=db.go -destination=mocks/mock_db.go -package=dbmocks

type Client interface {
	Group
	User
}

type client struct {
	Group
	User
}

func NewDB(cfg config.DatabaseConfig) (Client, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), gormconfig.GetGormConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	group := newGroup(db)
	user := newUser(db)

	return &client{
		group,
		user,
	}, nil
}
