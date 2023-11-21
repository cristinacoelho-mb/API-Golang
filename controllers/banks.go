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
// @Tags         Bank
// @Accept       json
// @Produce      json
// @Param        banks  path      string  true  "Banks ALL"
// @Success      200  {object}  models.bank
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /banks/ [get]

func ExibeTodosBanks(c *gin.Context) {
	var banks []models.Bank
	database.DB.Find(&banks)
	c.JSON(http.StatusOK, gin.H{
		"Todas os bancos cadastrados!": banks})

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Olá:": nome + ": " + "Seja benvindo ao nosso Bank!",
	})
}

// NewBank godoc
// @Summary updates account
// @Description creates Resource directory
// @Tags Bank
// @Param BankId path int true "BankID"
// @Param BankName path string true "BankName"
// @Param BankCode path int true "BankCode"
// @Param BankAgency path string true "BankAgency"
// @Param BankAccount path string true "BankAccount"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /bank [post]

func NewBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Não foi possǘel cadastrar o banco, verifique os dados digitados!": err.Error()})
		return
	}
	database.DB.Create(&bank)
	c.JSON(http.StatusOK, gin.H{
		"Banco cadastro com sucesso!": bank})
}

// ShowBankById godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         bank
// @Accept       json
// @Produce      json
// @Param        bankid  path      int  true  "Bank ID"
// @Success      200  {object}  models.bank
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bank/{id} [get]

func BuscaBankPorID(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Bank não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, bank)
}

// ShowBankById godoc
// @Summary      Show an bank
// @Description  get string by ID
// @Tags         Bank
// @Accept       json
// @Produce      json
// @Param        bankid  path      int  true  "Bank ID"
// @Success      200  {object}  models.bank
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bank/{id} [get]
func DeletaBank(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.Delete(&bank, id)
	c.JSON(http.StatusOK, gin.H{"data": "Bank deletado com sucesso!"})
}

// UpdateBank godoc
// @Summary updates bank
// @Description creates Resource directory
// @Tags Bank
// @Param bankid path int true "Bank Id"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /bank/{id} [patch]
func EditaBank(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.First(&bank, id)

	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: Atualização não realizada, verifique os dados digitados!": err.Error()})
		return
	}

	database.DB.Model(&bank).UpdateColumns(bank)
	c.JSON(http.StatusOK, gin.H{
		"Dados atualizados com sucesso!": bank})
}
