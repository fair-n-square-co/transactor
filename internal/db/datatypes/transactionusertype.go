package datatypes

import "database/sql/driver"

type TransactionUserType string

const (
	PAYER TransactionUserType = "payer"
	PAYEE TransactionUserType = "payee"
)

func (t *TransactionUserType) Scan(value interface{}) error {
	*t = TransactionUserType(value.([]byte))
	return nil
}

func (t TransactionUserType) Value() (driver.Value, error) {
	return string(t), nil
}
