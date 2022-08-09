package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tulip/backend/models"
)

func GetManualTransfers(c *gin.Context) {
	mt, err := models.GetManualTransfers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, mt)
}

func GetManualTransferById(c *gin.Context) {
	id := c.Param("id")
	mt, err := models.GetManualTransferById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manual Transfer not found"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, mt)
}
