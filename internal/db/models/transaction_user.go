package models

import (
	"github.com/fair-n-square-co/transactions/internal/db/datatypes"
	"github.com/fair-n-square-co/transactions/internal/db/models/base"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/money"
)

// TransactionUser represents a user in a transaction
// It creates a many to many relationship between users and transactions
// It also stores the amount the user is involved in the transaction
// and the type of user in the transaction. For example, payer or payee
//
// An example of a transaction user is:
// TransactionID: 1
// UserID: 1
// Amount: 100
// TransactionUserType: payer
// This means that user with ID 1 paid $100 in transaction with ID 1
//
// TransactionID: 1
// UserID: 1
// Amount: 50
// TransactionUserType: payee
// This means that user with ID 1 owes $50 in transaction with ID 1
//
// Notice that the transaction and user can appear multiple times in the table. So there are two logical unique keys
// TransactionID, UserID and TransactionUserType as a composite key
// OR
// another UUID column. We chose UUID to keep the table schema simple
//
// When we update a Transaction, we need to update all the TransactionUser table as well and the constraints are:
// 1. The sum of all the payer amounts should be equal to the transaction amount
// 2. The sum of all the payee amounts should be equal to the transaction amount
// Only then a database transaction can be considered successful
type TransactionUser struct {
	base.PrimaryKey
	base.DateTime
	base.SoftDeleteModel
	TransactionID       uuid.UUID                     `gorm:"type:uuid;not null"`
	UserID              uuid.UUID                     `gorm:"type:uuid;not null"`
	Amount              money.Money                   `gorm:"not null;embedded;embeddedPrefix:amount_"`
	TransactionUserType datatypes.TransactionUserType `gorm:"not null;default:'payer';type:Transaction_user_type"`
}
