package controllers

import (
	"API/database"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosCustomers(c *gin.Context) {
	var customers []models.Customer
	database.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{
		"Todos os clientes cadastrados!": customers})

}

func NewCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Não foi possível cadastrar o cliente, verifique os dados digitados!": err.Error()})
		return
	}
	database.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{
		"Cliente cadastrado com sucesso!": customer})
}

func BuscaCustomerPorID(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Customer não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func DeletaCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.Delete(&customer, id)
	c.JSON(http.StatusOK, gin.H{"data": "Customer deletado com sucesso!"})
}

func EditaCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Atualização não realizada, verifique os dados digitados!": err.Error()})
		return
	}

	database.DB.Model(&customer).UpdateColumns(customer)
	c.JSON(http.StatusOK, gin.H{
		"Atualização realizada com sucesso!": customer})
}

func BuscaPorCPF_CNPJ(c *gin.Context) {
	var customer models.Customer
	CPF_CNPJ := c.Param("CPF_CNPJ")
	database.DB.Where(&models.Customer{CPF_CNPJ: CPF_CNPJ}).First(&customer)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Customer não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, customer)
}
