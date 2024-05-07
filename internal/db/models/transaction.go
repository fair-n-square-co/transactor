package models

import (
	"github.com/fair-n-square-co/transactions/internal/db/datatypes"
	"github.com/fair-n-square-co/transactions/internal/db/models/base"
)

// Transaction represents a transaction in the system
// It stores the amount of the transaction, the description of the transaction
// All the user involved are stored in the TransactionUser table
// The type of transaction is also stored. For example, payment or settlement for semantic purposes
// TODO: Create a TransactionGroup table to store the groups involved in the transaction
type Transaction struct {
	base.PrimaryKey
	base.DateTime
	base.SoftDeleteModel
	Amount          datatypes.MoneyType       `gorm:"not null;embedded;embeddedPrefix:amount_"`
	Description     string                    `gorm:"not null"`
	Type            datatypes.TransactionType `gorm:"not null;default:payment;type:transaction_type"`
	CreatorID 		uint     
    UpdatorID	    uint     
	Creator 		User					  `gorm:"foreignKey:CreatorID"`
	LastUpdatedUser User					  `gorm:"foreignKey:UpdatorID"`
}
// reference https://stackoverflow.com/questions/77268541/gorm-two-foreign-keys-to-one-table