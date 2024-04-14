package models

import (
	"github.com/fair-n-square-co/transactions/internal/db/models/base"
)

type User struct {
	base.PrimaryKey
	base.DateTime
	base.SoftDeleteModel
	Email     string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string
	Phone     string  `gorm:"unique;type:varchar(100)"`
	Groups    []Group `gorm:"many2many:user_groups;"`
}
