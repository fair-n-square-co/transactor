package datatypes

import "database/sql/driver"

type TransactionType string

const (
	PAYMENT    TransactionType = "payment"
	SETTLEMENT TransactionType = "settlement"
)

func (t *TransactionType) Scan(value interface{}) error {
	*t = TransactionType(value.([]byte))
	return nil
}

func (t TransactionType) Value() (driver.Value, error) {
	return string(t), nil
}
