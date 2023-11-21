package controllers

import (
	"API/database"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccounts godoc
// @Summary      Show an accounts
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        accounts  path      string  true  "AccountS ALL"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/ [get]

func ExibeTodasAccounts(c *gin.Context) {
	var accounts []models.Account
	database.DB.Find(&accounts)
	c.JSON(http.StatusOK, gin.H{
		"Todas as contas cadastradas no Bank": accounts})

}

// NewAccount godoc
// @Summary updates account
// @Description creates Resource directory
// @Tags Accounts
// @Param AccountId path int true "Account Id"
// @Param FkCustomerId  path int true "Fk_Customer Id"
// @Param AccountName path string true "Account Name"
// @Param AccountNumber path string true "Account Number"
// @Param IsActive path bool true "Is Active"
// @Param AccountType path string true "Account Type"
// @Param CurrentBalance path string true "Current Balance"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account [post]

func NewAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Não foi possível cadastrar a conta, verifique se os dados estão corretos!": err.Error()})
		return
	}
	database.DB.Create(&account)
	c.JSON(http.StatusOK, gin.H{
		"Conta cadastrada com sucesso!": account})
}

// ShowAccountById godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        accoutsid  path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]

func BuscaAccountPorID(c *gin.Context) {
	var accountid models.Account
	id := c.Params.ByName("id")
	database.DB.First(&accountid, id)
	if accountid.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Account não encontrada!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Conta encontrada!": accountid})
}

// DeleteAccount godoc
// @Summary delete account
// @Description creates Resource directory
// @Tags Account
//@Param accountid  path string true "Account Id"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account/delete [delete]

func DeletaAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	database.DB.Delete(&account, id)
	c.JSON(http.StatusOK, gin.H{"data": "Account deletada com sucesso!"})
}

// UpdateAccount godoc
// @Summary updates account
// @Description creates Resource directory
// @Tags Accounts
// @Param accountid path string true "Account Id"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account/{id} [patch]
func EditaAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	database.DB.First(&account, id)

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Não foi possível atualizar os dados. Tente novamete!": err.Error()})
		return
	}
	database.DB.Model(&account).UpdateColumns(account)
	c.JSON(http.StatusOK, gin.H{
		"Conta atualizada com sucesso!": account})
}
