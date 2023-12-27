package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	BankId      int    `sql:"bank_id, type:int PRIMARY KEY"`
	BankName    string `sql:"bank_name, type:text"`
	BankCode    int64  `sql:"bank_code, type: int UNIQUE"`
	BankAgency  int64  `sql:"bank_agency, type:numeric"`
	BankAccount int64  `sql:"bank_account, type:numeric"`
}
