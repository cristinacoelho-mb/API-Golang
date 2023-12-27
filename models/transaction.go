package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionId         int64 `sql:"transaction_id,type:numeric PRIMARY KEY"`
	FkAccountId           int64 `sql:"fk_account_id, type:numeric REFERENCES account(account_id)"`
	WithdrawlValue        int64 `sql:"withdrawlValue, type:numeric(15,2)"`
	depositValue          int64 `sql:"depositValue, type:numeric(15,2)"`
	Transfer Value        int64 `sql:"transferValue, type:numeric(15,2)"`
	DestinationAccount    int64 `sql:"destinationAccount, type:numeric"`
	DestinationAgency     int64 `sql:"destinationAgency, type:text"`
	DestinationBank       int64 `sql:"destinationBank, type:text"`
	CurrentBalance        int64 `sql:"current_balance, type:numeric(15,2)"`
	
}
