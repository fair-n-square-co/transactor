package models

var models = []interface{}{
	// &datatypes.TransactionType{},
	&User{},
	&Transaction{},
	&Group{},
	&TransactionUser{},
}

func GetAllModels() []interface{} {
	return models
}
