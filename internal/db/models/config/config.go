package config

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var config = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
}

func GetGormConfig() *gorm.Config {
	return config
}

// Setup DB to use TransactionUser table as join table.
