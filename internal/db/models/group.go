package models

import "github.com/fair-n-square-co/transactions/internal/db/models/base"

// Group represents a group of users
type Group struct {
	base.PrimaryKey
	base.DateTime
	base.SoftDeleteModel
	Name  string `gorm:"not null"`
	Users []User `gorm:"many2many:user_groups;"`
}
