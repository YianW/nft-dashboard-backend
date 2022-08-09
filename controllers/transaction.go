package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tulip/backend/models"
)

func GetTransactions(c *gin.Context) {
	transactions, err := models.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, transactions)
}

func GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	t, err := models.GetTransactionById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, t)
}
