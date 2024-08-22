package models

import (
	"github.com/fair-n-square-co/transactions/internal/db/models/base"
)

// TODO : make change in proto to start accepting phone
// as it is going as empty string and has unique contraint over it so cannot add more than one record
type User struct {
	base.PrimaryKey
	base.DateTime
	base.SoftDeleteModel
	Email        string `gorm:"unique;not null"`
	Username     string `gorm:"unique;not null"`
	FirstName    string `gorm:"not null"`
	LastName     string
	Groups       []Group       `gorm:"many2many:user_groups;"`
	Transactions []Transaction `gorm:"many2many:transaction_user;"`
}
