ackage models

import (
	"gorm.io/gorm"
)

type dataCustomer struct {
	gorm.Model
	CustomerId       int    `sql:"customer_id, type:int PRIMARY KEY"`
	Name             string `json:"name" sql:"name, type:text NOT NULL"`
	CPF_CNPJ         string `json:"CPF OR CNPJ" sql:"CPF OR CNPJ, type: text NOT NULL"`
	PhoneNumber1     string `json:"phonenumber1" sql:"phonenumber1, type: text NOT NULL"`
	PhoneNumber2     string `json:"phonenumber2" sql:"phonenumber2, type: text NOT NULL"`
	Address          string `json:"address" sql:"address, type: text NOT NULL"`
	Email            string `json:"email" sql:"email, type:text NOT NULL"`
	
}