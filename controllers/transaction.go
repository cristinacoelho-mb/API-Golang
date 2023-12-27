package controllers

import (
	"API/database"
	"API/models"
	"fmt"
	"net/http"

	_ "gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExibeTodasTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"Todas as transactions": transactions})

}

func NewTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Não foi possível realizar está transaction!": err.Error()})
		return
	}
	database.DB.Create(&transaction)
	c.JSON(http.StatusOK, gin.H{
		"Transaction realizada com sucesso!": transaction})
}

func BuscaTransactionPorID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)
	if transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "transaction não encontrada!"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func DeletaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.Delete(&transaction, id)
	c.JSON(http.StatusOK, gin.H{"data": "Transaction deletado com sucesso!"})
}

func EditaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&transaction).UpdateColumns(transaction)
	c.JSON(http.StatusOK, gin.H{
		"Atualização realizada com sucesso!": transaction})
}

//corrigir regra de negocio



func Deposit(c *Account) (depositValue float64, destinationAccount int64 ) string{
	if depositValue > 0 {
	CurrentBalance += depositValue 
	return "Depósito realizado com sucesso"
} else {
	return "Depósito não realizado! Verifique os dados digitados!
  }





